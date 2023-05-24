package rserrors

import "net/http"

type UnAuthorizedError struct {
	Err  error
	Meta any
}

func (e UnAuthorizedError) RenderErrorResponse() (int, ErrorResponse) {
	return http.StatusUnauthorized, ErrorResponse{
		Errors: []ErrorObject{
			{
				Status: http.StatusUnauthorized,
				Title:  "Invalid token",
				Detail: e.Error(),
				Meta:   e.Meta,
			},
		},
	}
}

func (e UnAuthorizedError) Unwrap() error {
	return e.Err
}

func (e UnAuthorizedError) Error() string {
	return "Invalid token"
}
