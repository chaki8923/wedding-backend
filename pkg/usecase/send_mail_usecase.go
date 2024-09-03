package usecase

import (
	"net/smtp"
	"errors"
	"fmt"
	"github.com/chaki8923/wedding-backend/pkg/domain/model"
	"github.com/chaki8923/wedding-backend/pkg/domain/repository"
	// "github.com/chaki8923/wedding-backend/pkg/lib/config"
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

	// cfg, err := config.New()
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to load config: %w", err)
	// }

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

			recipientName := invitee.FamilyKj + " " + invitee.FirstKj
			// inviteeLink := "https://localhost:3443/invitee_detail?uuid=" + invitee.UUID + "&inv_id=" + *body
			// 本番
			inviteeLink := "https://front.wedding-hackathon.com/invitee_detail?uuid=" + invitee.UUID + "&inv_id=" + *body
			frontLink := "https://front.wedding-hackathon.com/sign_up"

			expiryDate := "2024-09-30" // 有効期限の日付（例）
			mailMessage := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s 様へ\r\n\r\nこの度、チャーキチャーキちゃんは結婚することになりました\r\n\r\n結婚式に招待したいのでまずは新規登録してください\r\n%s\r\n\r\nその後以下のリンクより出欠の更新を行ってください。\r\n%s\r\n\r\n有効期限: %s\r\n",
			recipient, *subject, recipientName, frontLink,inviteeLink, expiryDate)

			log.Printf("Sending email to: %s\n", recipient)

			// SMTPサーバ接続
			auth := smtp.PlainAuth("", "konkuriitonouenokareha128@gmail.com", "watobiceuwckrylb", "smtp.gmail.com")

			// メール送信
			err = smtp.SendMail("smtp.gmail.com:587", auth, "konkuriitonouenokareha128@gmail.com", []string{recipient}, []byte(mailMessage))
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