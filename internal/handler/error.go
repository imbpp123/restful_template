package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrorResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	StatusText string `json:"status"`
	ErrorText  string `json:"error,omitempty"`
}

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrorResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrorResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

func ErrNotFound() render.Renderer {
	return &ErrorResponse{
		HTTPStatusCode: 404,
		StatusText:     "Resource not found.",
	}
}

func ErrInternalError(err error) render.Renderer {
	return &ErrorResponse{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "Internal Server Error.",
		ErrorText:      err.Error(),
	}
}

func ErrByError(err error) render.Renderer {
	// logic of selecting errors from map or error
	// ...

	return &ErrorResponse{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "Internal Server Error.",
		ErrorText:      err.Error(),
	}
}
