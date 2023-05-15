package registry

import (
	"ReserveMate/backend/pkg/adapter/controller"
	"ReserveMate/backend/pkg/adapter/repository"
	"ReserveMate/backend/pkg/usecase/usecase"
)

func (r *registry) NewUserController() controller.User {
	u := usecase.NewUserUsecase(
		repository.NewUserRepository(r.db),
		repository.NewDBRepository(r.db),
	)

	return controller.NewUserController(u)
}
