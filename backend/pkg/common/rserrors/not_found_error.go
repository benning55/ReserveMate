package rserrors

import "net/http"

type NotFoundError struct {
	Err  error
	Meta any
}

func (e NotFoundError) RenderErrorResponse() (int, ErrorResponse) {
	return http.StatusNotFound, ErrorResponse{
		Errors: []ErrorObject{
			{
				Status: http.StatusNotFound,
				Title:  "Not Found",
				Detail: e.Unwrap().Error(),
				Meta:   e.Meta,
			},
		},
	}
}

func (e NotFoundError) Unwrap() error {
	return e.Err
}

func (e NotFoundError) Error() string {
	return "Not found error"
}
