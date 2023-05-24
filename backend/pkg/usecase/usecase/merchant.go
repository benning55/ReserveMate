package usecase

import (
	"ReserveMate/backend/pkg/domain/model"
	"ReserveMate/backend/pkg/usecase/repository"
)

type merchantUsecase struct {
	merchantRepository repository.MerchantRepository
	dBRepository       repository.DBRepository
}

type Merchant interface {
	List(u []*model.Merchant) ([]*model.Merchant, error)
}

func NewMerchantUsecase(r repository.MerchantRepository, d repository.DBRepository) Merchant {
	return &merchantUsecase{
		merchantRepository: r,
		dBRepository:       d,
	}
}

func (mu *merchantUsecase) List(u []*model.Merchant) ([]*model.Merchant, error) {
	u, err := mu.merchantRepository.FindAll(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}
