package controller

import (
	"github.com/benning55/go-clean-arch/pkg/domain/model"
	"github.com/benning55/go-clean-arch/pkg/usecase/usecase"
	"net/http"
)

type userController struct {
	userUsecase usecase.User
}

type User interface {
	GetUsers(ctx Context) error
}

func NewUserController(us usecase.User) User {
	return &userController{
		userUsecase: us,
	}
}

func (uc userController) GetUsers(ctx Context) error {
	var u []*model.User

	u, err := uc.userUsecase.List(u)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, u)
}
