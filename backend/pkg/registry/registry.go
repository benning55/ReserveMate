package registry

import (
	"ReserveMate/backend/pkg/adapter/controller"

	"gorm.io/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		User:     r.NewUserController(),
		Merchant: r.NewMerchantController(),
	}
}
