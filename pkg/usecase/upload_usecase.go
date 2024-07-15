package usecase

import (
	"fmt"
	"time"
	"context"
	"log"
	"github.com/99designs/gqlgen/graphql"
	"github.com/chaki8923/wedding-backend/pkg/domain/model"
	"github.com/chaki8923/wedding-backend/pkg/domain/repository"
)

type Upload interface {
	UploadFile(
		comment *string, 
		file_url *string) (*model.UploadImage, error)
	UploadFileToS3(ctx context.Context, file_url graphql.Upload) (string, error)
	GetImages()([]*model.UploadImage, error)

}

type UploadUseCase struct {
	updRepo repository.UploadImage
}

func NewUploadUseCase(updRepo repository.UploadImage) Upload {
	UploadUseCase := UploadUseCase{updRepo: updRepo}
	return &UploadUseCase
}

func (u *UploadUseCase) UploadFile(
	comment *string, 
	file_url *string) (*model.UploadImage, error) {
	now := time.Now().Format("2006-01-02 15:04:05")


	upload := model.UploadImage{
		Comment:      *comment,
		CreatedAt: now,
		UpdatedAt: now,
		FileURL:   *file_url, // 画像URLを保存
	}

	created, err := u.updRepo.UploadFile(&upload)
	if err != nil {
		return nil, fmt.Errorf("useCase UploadUseCase err: %w", err)
	}

	return created, nil
}


func (u *UploadUseCase) UploadFileToS3(ctx context.Context, file_url graphql.Upload) (string, error) {
	fileURL, err := u.updRepo.UploadFileToS3(ctx, file_url)
	if err != nil {
		return "", fmt.Errorf("usecase failed to upload file to S3: %w", err)
	}
	return fileURL, nil
}



func (u *UploadUseCase) GetImages() ([]*model.UploadImage, error) {
	images, err := u.updRepo.GetImages()
	if err != nil {
		return nil, fmt.Errorf("resolver 登録画像 err %w", err)
	}
	for _, msg := range images {
		log.Printf("InviteeUseCase一覧: %+v\n", *msg)
	}

	return images, nil
}

