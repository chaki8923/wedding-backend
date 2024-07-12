package repository

import (
	"github.com/chaki8923/wedding-backend/pkg/domain/model"
	"context"
	"github.com/99designs/gqlgen/graphql"
)

type Upload interface {
	UploadFileToS3(context.Context, graphql.Upload) (string, error)
	UploadFile(upload *model.Upload) (*model.Upload, error)
}
