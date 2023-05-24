package controller

import (
	"ReserveMate/backend/pkg/common/rserrors"
	"ReserveMate/backend/pkg/domain/model"
	"ReserveMate/backend/pkg/usecase/usecase"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userController struct {
	userUsecase usecase.User
}

type User interface {
	GetUsers(ctx Context) error
	CreateUser(ctx Context) error
	GetUserByEmail(ctx Context) error
	LogIn(ctx Context) error
}

func NewUserController(us usecase.User) User {
	return &userController{
		userUsecase: us,
	}
}

func (uc userController) GetUsers(ctx Context) error {
	fmt.Println("adapter-controller")
	var u []*model.User

	u, err := uc.userUsecase.List(u)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, u)
}

func (uc userController) CreateUser(ctx Context) error {
	var params model.User

	if err := ctx.Bind(&params); err != nil {
		return err
	}

	u, err := uc.userUsecase.Create(&params)

	if err != nil {
		status, response := rserrors.BadRequestError{
			Err: err,
		}.RenderErrorResponse()
		return ctx.JSON(status, response)
	}

	return ctx.JSON(http.StatusCreated, u)
}

func (uc userController) GetUserByEmail(ctx Context) error {
	email := ctx.Param("email")
	u, err := uc.userUsecase.GetUserByEmail(email)
	if err != nil {
		status, response := rserrors.NotFoundError{
			Err: err,
			Meta: map[string]interface{}{
				"email": email,
			},
		}.RenderErrorResponse()
		return ctx.JSON(status, response)
	}
	return ctx.JSON(http.StatusOK, u)
}

func (uc userController) LogIn(ctx Context) error {
	var params model.User

	if err := ctx.Bind(&params); err != nil {
		return err
	}

	token, err := uc.userUsecase.LogIn(params.Email)
	if err != nil {
		status, response := rserrors.NotFoundError{
			Err: err,
		}.RenderErrorResponse()
		return ctx.JSON(status, response)
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}
