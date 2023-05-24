package repository

import (
	"ReserveMate/backend/pkg/domain/model"
	"ReserveMate/backend/pkg/usecase/repository"

	"gorm.io/gorm"
)

type merchantRepository struct {
	db *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) repository.MerchantRepository {
	return &merchantRepository{db}
}

func (mr *merchantRepository) FindAll(u []*model.Merchant) ([]*model.Merchant, error) {
	err := mr.db.Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}
