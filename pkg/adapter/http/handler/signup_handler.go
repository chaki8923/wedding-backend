package handler

import (
	"fmt"
	"net/http"
	"golang.org/x/xerrors"

	"github.com/chaki8923/wedding-backend/pkg/usecase"
	"github.com/labstack/echo/v4"
)

type Signup interface {
	SignupHandler() echo.HandlerFunc
}

type SignupHandler struct {
	AuthUseCase usecase.Auth
}

func NewSignHandler(au usecase.Auth) Signup {
	SignupHandler := SignupHandler{
		AuthUseCase: au,
	}
	return &SignupHandler
}

func (l *SignupHandler) SignupHandler() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var fv = &usecase.FormValue{
			Email:    c.FormValue("email"),
			Password: c.FormValue("password"),
		}

		if err = c.Validate(fv); err != nil {
			return xerrors.Errorf("signup validate err: %w", err)
		}

		userId, err := l.AuthUseCase.Signup(c, fv)
		if err != nil {
			return fmt.Errorf("signup failed err: %w", err)
		}

		return c.JSON(http.StatusOK, echo.Map{
			"userId": userId,
		})
	}
}
