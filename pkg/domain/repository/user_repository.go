package repository

import (
	"context"

	"github.com/chaki8923/wedding-backend/pkg/domain/model/graph"
	"github.com/chaki8923/wedding-backend/pkg/domain/model"
)

type User interface {
	GetMapInIDs(ctx context.Context, ids []string) (map[string]*graph.User, error)
	GetUserById(id string) (*graph.User, error)
	CreateUser(email string, password string) (*model.User, error)
	GetUserByEmail(email string) (*graph.User, error)
	Encrypt(plain string) (string, error)
	Decrypt(encrypted string) (string, error)
}
