package infra

import (
	"github.com/chaki8923/wedding-backend/pkg/domain/model"
	"github.com/chaki8923/wedding-backend/pkg/domain/repository"
	"golang.org/x/xerrors"
	"gorm.io/gorm"
	"log"
)

type allergyRepository struct {
	db *gorm.DB
}

// repository内のallergyRepositoryのインターフェースに定義したらここで実装しないとエラー
func NewAllergyRepository(db *gorm.DB) repository.Allergy {

	return &allergyRepository{db}
}

func (i *allergyRepository) GetAllergy() ([]*model.Allergy, error) {
	var records []model.Allergy
	log.Printf("infra層----------------------")
	if result := i.db.Find(&records); result.Error != nil {
		return nil, xerrors.Errorf("repository  招待状取得 err %w", result.Error)
	}

	var res []*model.Allergy
	for _, record := range records {
		record := record
		res = append(res, &record)
	}
	for _, msg := range res {
		log.Printf("allergyInfra: %+v\n", *msg)
	}

	return res, nil
}


