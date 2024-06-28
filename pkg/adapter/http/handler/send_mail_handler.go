//一旦はっとく

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/aws/session"
// 	"github.com/aws/aws-sdk-go/service/ses"
// 	"github.com/labstack/echo/v4"
// )

// type EmailRequest struct {
// 	To      string `json:"to" validate:"required,email"`
// 	Subject string `json:"subject" validate:"required"`
// 	Body    string `json:"body" validate:"required"`
// }

// func main() {
// 	e := echo.New()

// 	e.POST("/send-email", sendEmail)

// 	e.Logger.Fatal(e.Start(":8080"))
// }

// func sendEmail(c echo.Context) error {
// 	req := new(EmailRequest)
// 	if err := c.Bind(req); err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
// 	}

// 	if err := c.Validate(req); err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request data"})
// 	}

// 	sess, err := session.NewSession(&aws.Config{
// 		Region: aws.String("us-west-2"), // 適切なリージョンに変更してください
// 	})
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create AWS session"})
// 	}

// 	svc := ses.New(sess)

// 	input := &ses.SendEmailInput
