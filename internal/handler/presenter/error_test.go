package presenter_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"app/internal/data"
	"app/internal/handler/presenter"
)

func TestErrorRender(t *testing.T) {
	type caseData struct {
		name       string
		err        error
		HTTPStatus int
		expected   map[string]interface{}
	}
	type testCase struct {
		apiVersion presenter.ApiVersion
		data       []caseData
	}

	testCases := map[string]testCase{
		"v1": {
			apiVersion: presenter.APIVersion1,
			data: []caseData{
				{
					name:       "article not found",
					err:        data.ErrArticleNotFound,
					HTTPStatus: http.StatusNotFound,
					expected: map[string]interface{}{
						"status": "Article not found.",
					},
				},
				{
					name:       "random error",
					err:        errors.New("random error"),
					HTTPStatus: http.StatusInternalServerError,
					expected: map[string]interface{}{
						"status": "Internal Server Error",
					},
				},
				{
					name:       "parameter not found",
					err:        data.ErrParameterNotFound,
					HTTPStatus: http.StatusBadRequest,
					expected: map[string]interface{}{
						"status": "Parameter not found.",
					},
				},
			},
		},
		"v2": {
			apiVersion: presenter.APIVersion2,
			data: []caseData{
				{
					name:       "article not found",
					err:        data.ErrArticleNotFound,
					HTTPStatus: http.StatusNotFound,
					expected: map[string]interface{}{
						"status": "Article not found.",
					},
				},
				{
					name:       "random error",
					err:        errors.New("random error"),
					HTTPStatus: http.StatusInternalServerError,
					expected: map[string]interface{}{
						"status": "Internal Server Error",
					},
				},
				{
					name:       "parameter not found",
					err:        data.ErrParameterNotFound,
					HTTPStatus: http.StatusBadRequest,
					expected: map[string]interface{}{
						"status": "Parameter not found.",
					},
				},
			},
		},
		"any": {
			apiVersion: "any",
			data: []caseData{
				{
					name:       "article not found",
					err:        data.ErrArticleNotFound,
					HTTPStatus: http.StatusNotFound,
					expected: map[string]interface{}{
						"status": "Article not found.",
					},
				},
				{
					name:       "random error",
					err:        errors.New("random error"),
					HTTPStatus: http.StatusInternalServerError,
					expected: map[string]interface{}{
						"status": "Internal Server Error",
					},
				},
				{
					name:       "parameter not found",
					err:        data.ErrParameterNotFound,
					HTTPStatus: http.StatusBadRequest,
					expected: map[string]interface{}{
						"status": "Parameter not found.",
					},
				},
			},
		},
	}

	for name, tc := range testCases {
		for _, data := range tc.data {
			t.Run(fmt.Sprintf("%s - %s", name, data.name), func(t *testing.T) {
				// arrange
				rec := httptest.NewRecorder()
				req, _ := http.NewRequest("", "", nil)
				req = req.WithContext(context.WithValue(req.Context(), presenter.ApiVersionCtx{}, tc.apiVersion))

				// act
				presenter.NewError().Render(rec, req, data.err)

				actual := make(map[string]interface{})
				assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &actual))

				// assert
				assert.Equal(t, data.HTTPStatus, rec.Code)
				assert.Equal(t, data.expected, actual)
			})
		}
	}
}
