package infra

import (
	"errors"
	"github.com/chaki8923/wedding-backend/pkg/domain/repository"
	"github.com/chaki8923/wedding-backend/pkg/domain/model"
	"gorm.io/gorm"
)

type SendMailRepository struct {
	db *gorm.DB
}

func NewSendMailRepository(db *gorm.DB) repository.SendMail {
	return &SendMailRepository{db}
}


func (r *SendMailRepository) SendMail(mail *model.SendMail) (*model.SendMail, error) {
	if mail.To == "" || mail.From == "" || mail.Subject == "" || mail.Body == "" {
		return nil, errors.New("one or more fields are empty")
	}

	return mail, nil
}