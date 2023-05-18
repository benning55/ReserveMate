package rserrors

type ErrorObject struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Meta   any    `json:"meta"`
}

type ErrorResponse struct {
	Errors []ErrorObject `json:"errors"`
}

type RSError interface {
	RenderErrorResponse() (int, ErrorResponse)
	Unwrap() error
	Error() string
}
