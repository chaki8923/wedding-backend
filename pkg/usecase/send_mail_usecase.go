package usecase

import (
	"net/smtp"
	"errors"
	"fmt"
	"github.com/chaki8923/wedding-backend/pkg/domain/model"
	"github.com/chaki8923/wedding-backend/pkg/domain/repository"
	"github.com/chaki8923/wedding-backend/pkg/lib/config"
)

type SendMail interface {
	SendMail(to *string, from *string, subject *string, body *string) (*model.SendMail, error)
}

type MailFormValue struct {
	To    string `json:"to" validate:"required,email"`
	From string `json:"from" validate:"required,email"`
	Subject string `json:"subject" validate:"required"`
	Body string `json:"body" validate:"required"`
}

type MailUseCase struct {
	mailRepo repository.SendMail
}

func NewMailUseCase(mailRepo repository.SendMail) SendMail {
	SendMailUseCase := MailUseCase{mailRepo: mailRepo}
	return &SendMailUseCase
}


func (m *MailUseCase) SendMail(to *string, from *string, subject *string, body *string) (*model.SendMail, error) {
	if to == nil || from == nil || subject == nil || body == nil {
		return nil, errors.New("one or more parameters are nil")
	}

	cfg, err := config.New()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	mailMessage := []byte("To: " + *to + "\r\n" +
		"Subject: " + *subject + "\r\n" +
		"\r\n" +
		*body + "\r\n")

	// SMTPサーバ接続
	auth := smtp.PlainAuth("", cfg.GoogleAccount, cfg.GoogleApiKey, "smtp.gmail.com")

	// メール送信
	err = smtp.SendMail("smtp.gmail.com:587", auth, cfg.GoogleAccount, []string{*from}, mailMessage)
	if err != nil {
		return nil, err
	}

	return &model.SendMail{
		To:      *to,
		From:    *from,
		Subject: *subject,
		Body:    *body,
	}, nil
}