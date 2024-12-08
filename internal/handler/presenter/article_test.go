package presenter_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"app/internal/data"
	"app/internal/handler/presenter"
)

func TestArticleRender(t *testing.T) {
	type testCase struct {
		apiVersion presenter.ApiVersion
		data       *data.Article
		expected   map[string]interface{}
	}

	articleUUID := uuid.New()
	testCases := map[string]testCase{
		"v1": {
			apiVersion: presenter.APIVersion1,
			data: &data.Article{
				ID:        articleUUID,
				Author:    "John Doe",
				Title:     "Test Article",
				Text:      "Test Content",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			expected: map[string]interface{}{
				"author": "John Doe",
				"uuid":   articleUUID.String(),
				"text":   "Test Content",
				"title":  "Test Article",
			},
		},
		"v2": {
			apiVersion: presenter.APIVersion2,
			data: &data.Article{
				ID:        articleUUID,
				Author:    "John Doe",
				Title:     "Test Article",
				Text:      "Test Content",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			expected: map[string]interface{}{
				"author": "John Doe",
				"uuid":   articleUUID.String(),
				"text":   "Test Content",
				"title":  "Test Article",
				"url":    fmt.Sprintf("http://localhost:3333/v2/?id=%s", articleUUID.String()),
			},
		},
		"default": {
			apiVersion: "any",
			data: &data.Article{
				ID:        articleUUID,
				Author:    "John Doe",
				Title:     "Test Article",
				Text:      "Test Content",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			expected: map[string]interface{}{
				"author": "John Doe",
				"uuid":   articleUUID.String(),
				"text":   "Test Content",
				"title":  "Test Article",
				"url":    fmt.Sprintf("http://localhost:3333/v2/?id=%s", articleUUID.String()),
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// arrange
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest("", "", nil)
			req = req.WithContext(context.WithValue(req.Context(), presenter.ApiVersionCtx{}, tc.apiVersion))

			// act
			presenter.NewArticleResponse().Render(rec, req, tc.data)

			actual := make(map[string]interface{})
			assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &actual))

			// assert
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
