package repository

import (
	"github.com/chaki8923/wedding-backend/pkg/domain/model"
)

type Allergy interface {
	GetAllergy() ([]*model.Allergy, error)
}
