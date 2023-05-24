package controller

import (
	"ReserveMate/backend/pkg/domain/model"
	"ReserveMate/backend/pkg/usecase/usecase"
	"fmt"
	"net/http"
)

type merchantController struct {
	merchantUsecase usecase.Merchant
}

type Merchant interface {
	GetMerchants(ctx Context) error
}

func NewMerchantController(us usecase.Merchant) Merchant {
	return &merchantController{
		merchantUsecase: us,
	}
}

func (mc merchantController) GetMerchants(ctx Context) error {
	fmt.Println("adapter-controller")
	var u []*model.Merchant

	u, err := mc.merchantUsecase.List(u)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, u)
}
