package repository

import (
	"github.com/chaki8923/wedding-backend/pkg/domain/model"
	"context"
	"github.com/99designs/gqlgen/graphql"
)

type Invitee interface {
	UploadFileToS3(context.Context, graphql.Upload) (string, error)
	GetInvitee() ([]*model.Invitee, error)
	CreateInvitee(invitee *model.Invitee) (*model.Invitee, error)
	UpdateInvitee(id string, invitee *model.Invitee) (*model.Invitee, error)
	ShowInvitee(id string) (*model.Invitee, error)
}
