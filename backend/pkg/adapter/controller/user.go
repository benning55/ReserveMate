package controller

import (
	"ReserveMate/backend/pkg/domain/model"
	"ReserveMate/backend/pkg/usecase/usecase"
	"fmt"
	"net/http"
)

type userController struct {
	userUsecase usecase.User
}

type User interface {
	GetUsers(ctx Context) error
	CreateUser(ctx Context) error
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
	fmt.Printf("*********")
	fmt.Printf("%v", u)

	return ctx.JSON(http.StatusOK, u)
}

func (uc userController) CreateUser(ctx Context) error {
	var params model.User

	if err := ctx.Bind(&params); err != nil {
		return err
	}

	u, err := uc.userUsecase.Create(&params)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, u)
}
