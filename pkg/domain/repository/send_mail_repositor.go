package repository

import (
	"github.com/chaki8923/wedding-backend/pkg/domain/model"
)

type SendMail interface {
	SendMail(mail *model.SendMail) (*model.SendMail, error)
	FindInviteeByEmail(email string) (*model.Invitee, error)
}
