package router

import (
	"ReserveMate/backend/pkg/adapter/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", func(context echo.Context) error { return c.User.GetUsers(context) })
	e.GET("/users/email/:email", func(context echo.Context) error {
		return c.User.GetUserByEmail(context)
	})
	e.POST("/users", func(context echo.Context) error {
		return c.User.CreateUser(context)
	})

	return e
}
