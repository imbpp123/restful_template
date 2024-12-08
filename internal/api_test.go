package internal_test

import (
	"app/internal"
	"app/internal/handler/presenter"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestApiArticleCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// arrange
	newArtile := map[string]interface{}{
		"author":  "test author",
		"title":   "test title",
		"content": "some random text here",
	}
	body, err := json.Marshal(newArtile)
	assert.NoError(t, err)

	handler := internal.InitializeArticleHandler()

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v2/articles", bytes.NewReader(body))
	req.Header.Add("Content-Type", "application/json")

	r := chi.NewRouter()
	r.With(presenter.ApiVersionMiddleware("v2")).Method("POST", "/v2/articles", handler.Create())

	// act
	r.ServeHTTP(rec, req)

	// assert
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

	expected := make(map[string]interface{})
	assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &expected))

	assert.Equal(t, expected["author"], newArtile["author"])
	assert.Equal(t, expected["title"], newArtile["title"])
	assert.Equal(t, expected["content"], newArtile["content"])
	assert.NotEmpty(t, expected["uuid"])
}
