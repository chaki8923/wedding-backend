package handler

import (
	"log"
	"fmt"
	"net/http"
	"github.com/chaki8923/wedding-backend/pkg/usecase"
	"github.com/labstack/echo/v4"
)

// ---------------------------------------
// SMTPサーバ情報取得
// ---------------------------------------


// ---------------------------------------
// メール送信
// param1. 送信するメールアドレス  string
// param2. 送信するメールの件名   string
// param3. 送信するメールの本文   string
// return: error
// ---------------------------------------


type SendMail interface {
	SendMailHandler() echo.HandlerFunc
}

type SendMailHandler struct {
	SendMailUseCase usecase.SendMail
}

func NewSendMailHandler(su usecase.SendMail) SendMail {
	SendMailHandler := SendMailHandler{
		SendMailUseCase: su,
	}
	return &SendMailHandler
}

func (s *SendMailHandler) SendMailHandler() echo.HandlerFunc {
	log.Printf("mailhandler入った-------------------------")

	return func(c echo.Context) (err error) {
		var fv = &usecase.MailFormValue{
			To:    c.FormValue("to"),
			From: c.FormValue("from"),
			Subject: c.FormValue("subject"),
			Body: c.FormValue("body"),
		}


		userId, err := s.SendMailUseCase.SendMail(&fv.From, &fv.To, &fv.Subject, &fv.Body)
		if err != nil {
			return fmt.Errorf("signup failed err: %w", err)
		}

		return c.JSON(http.StatusOK, echo.Map{
			"userId": userId,
		})
	}

}
