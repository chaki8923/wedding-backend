package usecase

import (
	"fmt"
	"time"
	"log"
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/chaki8923/wedding-backend/pkg/domain/model"
	"github.com/chaki8923/wedding-backend/pkg/domain/repository"
)

type Invitee interface {
	CreateInvitee(
		family_kj *string, 
		first_kj *string, 
		family_kn *string, 
		first_kn *string, 
		email *string, 
		zip_code *string, 
		address_text *string, 
		allergy *string, 
		userId *string,
		file_url *string) (*model.Invitee, error)

	UpdateInvitee(
		id *string, 
		family_kj *string, 
		first_kj *string, 
		family_kn *string, 
		first_kn *string, 
		email *string, 
		zip_code *string, 
		address_text *string, 
		allergy *string, 
		file_url *string) (*model.Invitee, error)
	GetInvitee() ([]*model.Invitee, error)
	ShowInvitee(id string) (*model.Invitee, error)
	UploadFileToS3(ctx context.Context, file_url graphql.Upload) (string, error)

}

type IvteeUseCase struct {
	ivteeRepo repository.Invitee
}

func NewIvteeUseCase(ivteeRepo repository.Invitee) Invitee {
	InviteeUseCase := IvteeUseCase{ivteeRepo: ivteeRepo}
	return &InviteeUseCase
}

func (i *IvteeUseCase) CreateInvitee(
	family_kj *string, 
	first_kj *string, 
	family_kn *string, 
	first_kn *string, 
	email *string, 
	zip_code *string, 
	address_text *string, 
	userId *string,
	allergy *string, 
	file_url *string) (*model.Invitee, error) {
	now := time.Now().Format("2006-01-02 15:04:05")


	invitee := model.Invitee{
		FamilyKj:      *family_kj,
		FirstKj:  *first_kj,
		FamilyKn:      *family_kn,
		FirstKn:      *first_kn,
		Email:      *email,
		ZipCode:      *zip_code,
		AddressText:      *address_text,
		Allergy:      *allergy,
		UserID:    *userId,
		CreatedAt: now,
		UpdatedAt: now,
		FileURL:   *file_url, // 画像URLを保存
	}

	created, err := i.ivteeRepo.CreateInvitee(&invitee)
	if err != nil {
		return nil, fmt.Errorf("useCase CreateInvitation err: %w", err)
	}

	return created, nil
}

func (i *IvteeUseCase) UpdateInvitee(
	id *string, 
	family_kj *string, 
	first_kj *string, 
	family_kn *string, 
	first_kn *string, 
	email *string, 
	zip_code *string, 
	address_text *string, 
	allergy *string, 
	file_url *string) (*model.Invitee, error)  {
	now := time.Now().Format("2006-01-02 15:04:05")
	updatedInvitee := model.Invitee{
		FamilyKj:      *family_kj,
		FirstKj:  *first_kj,
		FamilyKn:      *family_kn,
		FirstKn:      *first_kn,
		Email:      *email,
		ZipCode:      *zip_code,
		AddressText:      *address_text,
		Allergy:      *allergy,
		UpdatedAt: now,
		FileURL:   *file_url, // 画像URLを保存
	}

	invitee, err := i.ivteeRepo.UpdateInvitee(*id, &updatedInvitee)
	if err != nil {
		return nil, fmt.Errorf("useCase UpdateInvitation err %w", err)
	}

	return invitee, nil
}

func (i *IvteeUseCase) GetInvitee() ([]*model.Invitee, error) {
	invitee, err := i.ivteeRepo.GetInvitee()
	if err != nil {
		return nil, fmt.Errorf("resolver 招待者 err %w", err)
	}
	for _, msg := range invitee {
		log.Printf("InviteeUseCase一覧: %+v\n", *msg)
	}

	return invitee, nil
}

func (i *IvteeUseCase) ShowInvitee(id string) (*model.Invitee, error) {
	invitee, err := i.ivteeRepo.ShowInvitee(id)
	if err != nil {
		return nil, fmt.Errorf("resolver 招待状 err %w", err)
	}


	log.Printf("InvitationUseCase詳細！！: %+v\n", invitee)

	return invitee, nil
}


func (i *IvteeUseCase) UploadFileToS3(ctx context.Context, file_url graphql.Upload) (string, error) {
	fileURL, err := i.ivteeRepo.UploadFileToS3(ctx, file_url)
	if err != nil {
		return "", fmt.Errorf("usecase failed to upload file to S3: %w", err)
	}
	return fileURL, nil
}


