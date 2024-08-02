package usecase

import (
	"fmt"
	"log"
	"github.com/chaki8923/wedding-backend/pkg/domain/model"
	"github.com/chaki8923/wedding-backend/pkg/domain/repository"
)

type Allergy interface {
	GetAllergy() ([]*model.Allergy, error)

}

type AgyUseCase struct {
	agyRepo repository.Allergy
}

func NewAgyUseCase(agyRepo repository.Allergy) Allergy {
	AllergyUseCase := AgyUseCase{agyRepo: agyRepo}
	return &AllergyUseCase
}



func (a *AgyUseCase) GetAllergy() ([]*model.Allergy, error) {
	allergies, err := a.agyRepo.GetAllergy()
	if err != nil {
		return nil, fmt.Errorf("resolver 招待状 err %w", err)
	}
	for _, msg := range allergies {
		log.Printf("InvitationUseCase一覧: %+v\n", *msg)
	}

	return allergies, nil
}

