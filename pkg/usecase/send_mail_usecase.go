package usecase

import (
	"net/smtp"
	"errors"
	"fmt"
	"github.com/chaki8923/wedding-backend/pkg/domain/model"
	"github.com/chaki8923/wedding-backend/pkg/domain/repository"
	"github.com/chaki8923/wedding-backend/pkg/lib/config"
	"strings"
	"log"
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

	// 複数の宛先に対応するためにカンマで分割
	recipients := strings.Split(*to, ",")
	log.Printf("送信先: %+v\n", recipients)
	for _, recipient := range recipients {
			recipient = strings.TrimSpace(recipient)
			if recipient == "" {
					continue
			}

			// ここで invitee を取得
			invitee, err := m.mailRepo.FindInviteeByEmail(recipient)
			if err != nil {
				return nil, fmt.Errorf("failed to find invitee with email %s: %w", recipient, err)
			}

			log.Printf("招待者情報: %+v\n", invitee)
			recipientName := invitee.FamilyKj + " " + invitee.FirstKj
		  customBody := fmt.Sprintf("%s 様へ\n\n%s", recipientName, *body)

			mailMessage := []byte("To: " + recipient + "\r\n" +
					"Subject: " + *subject + "\r\n" +
					"\r\n" +
					customBody + "\r\n")
				log.Printf("Sending email to: %s\n", recipient)
			// SMTPサーバ接続
			auth := smtp.PlainAuth("", cfg.GoogleAccount, cfg.GoogleApiKey, "smtp.gmail.com")

			// メール送信
			err = smtp.SendMail("smtp.gmail.com:587", auth, cfg.GoogleAccount, []string{recipient}, mailMessage)
			if err != nil {
				return nil, fmt.Errorf("failed to send email to %s: %w", recipient, err)
			}
	}

	return &model.SendMail{
		To:      *to,
		From:    *from,
		Subject: *subject,
		Body:    *body,
	}, nil
}