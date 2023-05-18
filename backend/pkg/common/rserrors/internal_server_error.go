package rserrors

import (
	"net/http"
)

type InternalServerError struct {
	Err error
}

func (e InternalServerError) RenderErrorResponse() (int, ErrorResponse) {
	return http.StatusInternalServerError, ErrorResponse{
		Errors: []ErrorObject{
			{
				Status: http.StatusInternalServerError,
				Title:  "Internal Server Error",
				Detail: e.Error(),
				Meta:   make(map[string]interface{}),
			},
		},
	}
}

func (e InternalServerError) Unwrap() error {
	return e.Err
}

func (e InternalServerError) Error() string {
	return "Internal Server Error"
}
