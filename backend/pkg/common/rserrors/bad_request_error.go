package rserrors

import "net/http"

type BadRequestError struct {
	Err  error
	Meta any
}

func (e BadRequestError) RenderErrorResponse() (int, ErrorResponse) {
	return http.StatusBadRequest, ErrorResponse{
		Errors: []ErrorObject{
			{
				Status: http.StatusBadRequest,
				Title:  "Bad request",
				Detail: e.Unwrap().Error(),
				Meta:   e.Meta,
			},
		},
	}
}

func (e BadRequestError) Unwrap() error {
	return e.Err
}

func (e BadRequestError) Error() string {
	return "Bad request"
}
