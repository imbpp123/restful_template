package presenter

import (
	"net/http"

	"github.com/go-chi/render"

	v1 "app/internal/handler/presenter/v1"
	v2 "app/internal/handler/presenter/v2"
)

type Error struct {
}

func NewError() *Error {
	return &Error{}
}

func (e *Error) Render(w http.ResponseWriter, r *http.Request, err error) {
	var payload render.Renderer

	switch getAPIVersion(r) {
	case APIVersion1:
		payload = v1.NewErrorResponse(err)
	case APIVersion2:
		payload = v2.NewErrorResponse(err)
	default:
		payload = v2.NewErrorResponse(err)
	}

	render.Render(w, r, payload)
}
