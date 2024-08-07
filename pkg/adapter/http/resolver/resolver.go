package resolver

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/chaki8923/wedding-backend/pkg/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	MsgUseCase  usecase.Message
	UserUseCase usecase.User
	IvtUseCase usecase.Invitation
	IvteeUseCase usecase.Invitee
	UpdUseCase usecase.Upload
	AgyUseCase usecase.Allergy
}
