package repository

import (
	"github.com/chaki8923/wedding-backend/pkg/domain/model"
)

type Message interface {
	GetMessages() ([]*model.Message, error)
	CreateMessage(todos *model.Message) (*model.Message, error)
}
