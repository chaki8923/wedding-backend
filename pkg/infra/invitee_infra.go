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
	"time"
	"github.com/chaki8923/wedding-backend/pkg/lib/config"
	"os"
	"io"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

type inviteeRepository struct {
	db *gorm.DB
	s3Client *s3.S3
	bucket   string
}

// repository内のinvitationRepositoryのインターフェースに定義したらここで実装しないとエラー
func NewInviteeRepository(db *gorm.DB, s3Client *s3.S3, bucket string) repository.Invitee {
	return &inviteeRepository{
		db: db,
		s3Client: s3Client,
		bucket: bucket,
	}
}

func (i *inviteeRepository) GetInvitee() ([]*model.Invitee, error) {
	var records []model.Invitee
	log.Printf("infra層---------------------")
	if result := i.db.Find(&records); result.Error != nil {
		return nil, xerrors.Errorf("repository  招待者取得 err %w", result.Error)
	}

	var res []*model.Invitee
	for _, record := range records {
		record := record
		res = append(res, &record)
	}
	for _, msg := range res {
		log.Printf("InviteeInfra: %+v\n", *msg)
	}

	return res, nil
}

func (i *inviteeRepository) CreateInvitee(invitee *model.Invitee) (*model.Invitee, error) {

	if result := i.db.Create(invitee); result.Error != nil {
		return nil, xerrors.Errorf("repository 招待者取得 err %w", result.Error)
	}
	return invitee, nil
}

func (i *inviteeRepository) UpdateInvitee(id string, updatedInvitee *model.Invitee) (*model.Invitee, error) {
	// まず、対象のレコードを取得する
	var invitee model.Invitee
	if result := i.db.First(&invitee, "id = ?", id); result.Error != nil {
		return nil, xerrors.Errorf("repository 招待状更新Infra err %w", result.Error)
	}
	log.Printf("infra_join!!!: %+v\n", updatedInvitee.JoinFlag)


	// レコードを更新する前に現在の状態をログに出力
	log.Printf("Before update: %+v\n", invitee)
	// 更新内容をマップに設定
	updates := map[string]interface{}{
		"UpdatedAt": time.Now().Format("2006-01-02 15:04:05"),
	}

	if updatedInvitee.FamilyKj != "" {
			updates["FamilyKj"] = updatedInvitee.FamilyKj
	}
	if updatedInvitee.FirstKj != "" {
			updates["FirstKj"] = updatedInvitee.FirstKj
	}
	if updatedInvitee.FamilyKn != "" {
		updates["FamilyKn"] = updatedInvitee.FamilyKn
		}
	if updatedInvitee.FirstKn != "" {
			updates["FirstKn"] = updatedInvitee.FirstKn
	}
	if updatedInvitee.Email != "" {
			updates["Email"] = updatedInvitee.Email
	}
	if updatedInvitee.ZipCode != "" {
			updates["ZipCode"] = updatedInvitee.ZipCode
	}
	if updatedInvitee.FileURL != "" {
			updates["FileURL"] = updatedInvitee.FileURL
	}
	if updatedInvitee.AddressText != "" {
			updates["AddressText"] = updatedInvitee.AddressText
	}
	if updatedInvitee.Allergy != "" {
			updates["Allergy"] = updatedInvitee.Allergy
	}

	// boolean型のフィールドはnilチェックではなくそのまま設定
	updates["JoinFlag"] = updatedInvitee.JoinFlag

	// 更新後の状態をログに出力
	log.Printf("After update: %+v\n", updates)

	if result := i.db.Model(&invitee).Omit("created_at").Updates(updates); result.Error != nil {
		return nil, xerrors.Errorf("repository 招待状更新 err %w", result.Error)
	}



	return &invitee, nil
}

func (i *inviteeRepository) ShowInvitee(id string) (*model.Invitee, error) {
	var record model.Invitee
	if result := i.db.Where("id = ?", id).First(&record); result.Error != nil {
		return nil, xerrors.Errorf("repository  招待状詳細取得 err %w", result.Error)
	}

	log.Printf("InvitationInfra詳細: %+v\n", record)

	return &record, nil
}

func (i *inviteeRepository) DeleteInvitee(id string) (*model.Invitee, error) {
	var record model.Invitee
	if result := i.db.Where("id = ?", id).First(&record); result.Error != nil {
		return nil, xerrors.Errorf("repository  招待状詳細取得 err %w", result.Error)
	}

	  // S3の画像を削除
    if record.FileURL != "" {
			_, err := i.s3Client.DeleteObject(&s3.DeleteObjectInput{
					Bucket: aws.String(i.bucket),
					Key:    aws.String(record.FileURL),
			})
			if err != nil {
					return nil, xerrors.Errorf("S3画像削除 err %w", err)
			}

			err = i.s3Client.WaitUntilObjectNotExists(&s3.HeadObjectInput{
					Bucket: aws.String(i.bucket),
					Key:    aws.String(record.FileURL),
			})
			if err != nil {
					return nil, xerrors.Errorf("S3画像削除確認 err %w", err)
			}
	}

	if result := i.db.Delete(&record); result.Error != nil {
		return nil, xerrors.Errorf("repository 招待状削除 err %w", result.Error)
}

	log.Printf("InvitationInfra詳細: %+v\n", record)

	return &record, nil
}


func (i *inviteeRepository) UploadFileToS3(ctx context.Context, file_url graphql.Upload) (string, error) {
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

	fileKey := "invitee/" + file_url.Filename
	uploadInput := &s3.PutObjectInput{
		Bucket: aws.String("weddingnet"),
		Key:    aws.String(fileKey),
		Body:   tempFile,
	}
	_, err = uploader.PutObjectWithContext(ctx, uploadInput)
	if err != nil {
		return "", fmt.Errorf("failed to upload file to S3: %w", err)
	}

	// ファイルのアップロード後の URL を返す
	fileUrl := "https://weddingnet.s3-ap-northeast-1.amazonaws.com/" + fileKey
	log.Printf("file_url %s", fileUrl)
	return fileUrl, nil
}