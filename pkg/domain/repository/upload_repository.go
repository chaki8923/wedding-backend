package repository

import (
	"github.com/chaki8923/wedding-backend/pkg/domain/model"
	"context"
	"github.com/99designs/gqlgen/graphql"
)

type UploadImage interface {
	UploadFileToS3(context.Context, graphql.Upload) (string, error)
	UploadFile(upload *model.UploadImage) (*model.UploadImage, error)
	GetImages() ([]*model.UploadImage, error)
}
