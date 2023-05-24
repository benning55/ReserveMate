package registry

import (
	"ReserveMate/backend/pkg/adapter/controller"
	"ReserveMate/backend/pkg/adapter/repository"
	"ReserveMate/backend/pkg/usecase/usecase"
)

func (r *registry) NewMerchantController() controller.Merchant {
	u := usecase.NewMerchantUsecase(
		repository.NewMerchantRepository(r.db),
		repository.NewDBRepository(r.db),
	)

	return controller.NewMerchantController(u)
}
