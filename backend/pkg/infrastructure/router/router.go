package router

import (
	"ReserveMate/backend/pkg/adapter/controller"
	"ReserveMate/backend/pkg/common/rserrors"
	"ReserveMate/backend/pkg/middlewares"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("v1/login", func(context echo.Context) error { return c.User.LogIn(context) })

	r := e.Group("/v1")
	config := echojwt.Config{
		NewClaimsFunc: func(context echo.Context) jwt.Claims {
			return new(middlewares.JwtCustomClaims)
		},
		SigningKey:  []byte("secret"),
		TokenLookup: "header:Authorization:Bearer ",
		ErrorHandler: func(ctx echo.Context, err error) error {
			status, response := rserrors.UnAuthorizedError{
				Err: err,
			}.RenderErrorResponse()
			return ctx.JSON(status, response)
		},
	}
	r.Use(echojwt.WithConfig(config))

	r.GET("/users", func(context echo.Context) error { return c.User.GetUsers(context) })
	r.GET("/users/email/:email", func(context echo.Context) error {
		return c.User.GetUserByEmail(context)
	})
	r.POST("/users", func(context echo.Context) error {
		return c.User.CreateUser(context)
	})

	r.GET("/merchants", func(context echo.Context) error { return c.Merchant.GetMerchants(context) })

	return e
}
