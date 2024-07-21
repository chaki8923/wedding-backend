package repository

import (
	"github.com/chaki8923/wedding-backend/pkg/domain/model"
	"context"
	"github.com/99designs/gqlgen/graphql"
)

type Invitation interface {
	UploadFileToS3(context.Context, graphql.Upload) (string, error)
	GetInvitation() ([]*model.Invitation, error)
	CreateInvitation(invitation *model.Invitation) (*model.Invitation, error)
	UpdateInvitation(id string, invitation *model.Invitation) (*model.Invitation, error)
	ShowInvitation(uu_id string) (*model.Invitation, error)
	DeleteInvitation(id string) (*model.Invitation, error)
}
