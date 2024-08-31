package infra

import (
	"github.com/chaki8923/wedding-backend/pkg/domain/model"
	"github.com/99designs/gqlgen/graphql"
	"github.com/chaki8923/wedding-backend/pkg/domain/repository"
	"golang.org/x/xerrors"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"gorm.io/gorm"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"github.com/chaki8923/wedding-backend/pkg/lib/config"
	"os"
	"io"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

type updRepository struct {
	db *gorm.DB
	s3Client *s3.S3
	bucket   string
}

// repository内のinvitationRepositoryのインターフェースに定義したらここで実装しないとエラー
func NewUploadRepository(db *gorm.DB, s3Client *s3.S3, bucket string) repository.UploadImage {
	return &updRepository{
		db: db,
		s3Client: s3Client,
		bucket: bucket,
	}
}

func (u *updRepository) UploadFile(upload *model.UploadImage) (*model.UploadImage, error) {

	if result := u.db.Create(upload); result.Error != nil {
		return nil, xerrors.Errorf("repository 招待者取得 err %w", result.Error)
	}
	return upload, nil
}

func (u *updRepository) UploadFileToS3(ctx context.Context, file_url graphql.Upload) (string, error) {
	// ファイルの一時保存先の準備
	tempFile, err := os.CreateTemp("", "upload-*.tmp")
	if err != nil {
		return "", fmt.Errorf("failed to create temporary file: %w", err)
	}

	cfg, err := config.New()
	if err != nil {
		return "", fmt.Errorf("failed to load config: %w", err)
	}

	// ファイルを一時ファイルに書き込む
	written, err := io.Copy(tempFile, file_url.File)
	if err != nil {
		return "", fmt.Errorf("failed to write to temporary file: %w", err)
	}

	// 書き込まれたバイト数をログ出力して確認
	log.Printf("written %d bytes to temp file", written)
	// 一時ファイルを再オープンしてポインタを先頭に戻す
	tempFile, err = os.Open(tempFile.Name())
	if err != nil {
		return "", fmt.Errorf("failed to reopen temporary file: %w", err)
	}

	defer tempFile.Close()

	credential := credentials.NewStaticCredentials(
		cfg.AwsAccessKey,
		cfg.AwsSecretKey,
		"",
	)

	// ファイルを S3 にアップロードする処理
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Credentials: credential,
	}))
	uploader := s3.New(sess)

	fileKey := "upload_file/" + file_url.Filename
	uploadInput := &s3.PutObjectInput{
		// Bucket: aws.String("weddingnet"),
		Bucket: aws.String("wedding-gate"),
		Key:    aws.String(fileKey),
		Body:   tempFile,
	}
	_, err = uploader.PutObjectWithContext(ctx, uploadInput)
	if err != nil {
		return "", fmt.Errorf("failed to upload file to S3: %w", err)
	}

	// ファイルのアップロード後の URL を返す
	// fileUrl := "https://weddingnet.s3-ap-northeast-1.amazonaws.com/" + fileKey
	fileUrl := "https://wedding-gate.s3-ap-northeast-1.amazonaws.com/" + fileKey
	log.Printf("file_url %s", fileUrl)
	return fileUrl, nil
}

func (u *updRepository) GetImages() ([]*model.UploadImage, error) {
	var records []model.UploadImage
	if result := u.db.Find(&records); result.Error != nil {
		return nil, xerrors.Errorf("repository  招待者取得 err %w", result.Error)
	}

	var res []*model.UploadImage
	for _, record := range records {
		record := record
		res = append(res, &record)
	}
	for _, msg := range res {
		log.Printf("画像等: %+v\n", *msg)
	}

	return res, nil
}
