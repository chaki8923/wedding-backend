package usecase

import (
	"fmt"
	"time"
	"log"
	"context"
	"github.com/google/uuid"
	"github.com/99designs/gqlgen/graphql"
	"github.com/chaki8923/wedding-backend/pkg/domain/model"
	"github.com/chaki8923/wedding-backend/pkg/domain/repository"
)

type Invitation interface {
	CreateInvitation(title *string, event_date *string, place *string, comment *string, userId *string, file_url *string, uu_id *string) (*model.Invitation, error)
	UpdateInvitation(id *string, title *string, event_date *string, place *string, comment *string) (*model.Invitation, error)
	GetInvitation() ([]*model.Invitation, error)
	ShowInvitation(id string) (*model.Invitation, error)
	DeleteInvitation(id string) (*model.Invitation, error)
	UploadFileToS3(ctx context.Context, file_url graphql.Upload) (string, error)

}

type IvtUseCase struct {
	ivtRepo repository.Invitation
}

func NewIvtUseCase(ivtRepo repository.Invitation) Invitation {
	InvitationUseCase := IvtUseCase{ivtRepo: ivtRepo}
	return &InvitationUseCase
}

func (i *IvtUseCase) CreateInvitation(title *string, event_date *string, place *string, comment *string, userId *string, file_url *string, uu_id *string) (*model.Invitation, error) {
	now := time.Now().Format("2006-01-02 15:04:05")
	// UUIDを生成
	newUUID := uuid.New().String()

	invitation := model.Invitation{
		Title:      *title,
		EventDate:  *event_date,
		Place:      *place,
		Comment:      *comment,
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    *userId,
		FileURL:   *file_url, // 画像URLを保存
		UUID:       newUUID, // 生成したUUIDを設定
	}

	created, err := i.ivtRepo.CreateInvitation(&invitation)
	if err != nil {
		return nil, fmt.Errorf("useCase CreateInvitation err: %w", err)
	}

	return created, nil
}

func (i *IvtUseCase) UpdateInvitation(id *string, title *string, eventDate *string, place *string, comment *string) (*model.Invitation, error) {
	now := time.Now().Format("2006-01-02 15:04:05")
	updatedInvitation := model.Invitation{
		Title:     *title,
		EventDate: *eventDate,
		Place:     *place,
		Comment:     *comment,
		UpdatedAt: now,
	}

	invitation, err := i.ivtRepo.UpdateInvitation(*id, &updatedInvitation)
	if err != nil {
		return nil, fmt.Errorf("useCase UpdateInvitation err %w", err)
	}

	return invitation, nil
}

func (i *IvtUseCase) GetInvitation() ([]*model.Invitation, error) {
	invitations, err := i.ivtRepo.GetInvitation()
	if err != nil {
		return nil, fmt.Errorf("resolver 招待状 err %w", err)
	}
	for _, msg := range invitations {
		log.Printf("InvitationUseCase一覧: %+v\n", *msg)
	}

	return invitations, nil
}

func (i *IvtUseCase) ShowInvitation(uu_id string) (*model.Invitation, error) {
	invitation, err := i.ivtRepo.ShowInvitation(uu_id)
	if err != nil {
		return nil, fmt.Errorf("resolver 招待状 err %w", err)
	}


	log.Printf("InvitationUseCase詳細！！: %+v\n", invitation)

	return invitation, nil
}


func (i *IvtUseCase) UploadFileToS3(ctx context.Context, file_url graphql.Upload) (string, error) {
	fileURL, err := i.ivtRepo.UploadFileToS3(ctx, file_url)
	if err != nil {
		return "", fmt.Errorf("usecase failed to upload file to S3: %w", err)
	}
	return fileURL, nil
}

func (i *IvtUseCase) DeleteInvitation(id string) (*model.Invitation, error) {
	invitation, err := i.ivtRepo.DeleteInvitation(id)
	if err != nil {
		return nil, fmt.Errorf("resolver 招待状 err %w", err)
	}


	log.Printf("InvitationUseCase駆除！！: %+v\n", invitation)

	return invitation, nil
}
