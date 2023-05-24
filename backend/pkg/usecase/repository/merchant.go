package repository

import "ReserveMate/backend/pkg/domain/model"

type MerchantRepository interface {
	FindAll(u []*model.Merchant) ([]*model.Merchant, error)
}
