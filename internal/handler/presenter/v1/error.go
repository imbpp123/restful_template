package v1

import (
	"app/internal/data"
	"net/http"

	"github.com/go-chi/render"
)

type ErrorResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	StatusText string `json:"status"`
	AppCode    int64  `json:"code,omitempty"`
	ErrorText  string `json:"error,omitempty"`
}

var (
	errorResponseList = map[error]ErrorResponse{
		data.ErrArticleNotFound: {
			HTTPStatusCode: http.StatusNotFound,
			StatusText:     "Article not found.",
		},
		data.ErrParameterNotFound: {
			HTTPStatusCode: http.StatusBadRequest,
			StatusText:     "Parameter not found.",
		},
	}

	defaultErrorResponse = ErrorResponse{
		HTTPStatusCode: http.StatusInternalServerError,
		StatusText:     "Internal Server Error",
	}
)

func NewErrorResponse(err error) *ErrorResponse {
	if response, ok := errorResponseList[err]; ok {
		return &response
	}

	return &defaultErrorResponse
}

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}
