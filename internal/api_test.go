package internal_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"app/internal"
	"app/internal/data"
	"app/internal/domain"
	"app/internal/handler"
	"app/internal/repository"
)

type articleAPIUnitTest struct {
	handler    *handler.ArticleHandler
	repository *repository.Article
}

func newArticleAPIUnitTest() *articleAPIUnitTest {
	articleRepository := repository.NewArticle()
	articleService := domain.NewArticle(
		articleRepository,
	)

	return &articleAPIUnitTest{
		handler: handler.NewArticleHandler(
			articleService,
			validator.New(),
		),
		repository: articleRepository,
	}
}

func TestApiArticleCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// arrange
	newArtile := map[string]interface{}{
		"author": "test author",
		"title":  "test title",
		"text":   "some random text here",
	}
	body, err := json.Marshal(newArtile)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/articles", bytes.NewReader(body))
	req.Header.Add("Content-Type", "application/json")

	ut := newArticleAPIUnitTest()

	articleHandler := ut.handler
	r := internal.RouterInitializer(articleHandler)

	// act
	r.ServeHTTP(rec, req)

	// assert
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

	expected := make(map[string]interface{})
	assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &expected))

	assert.Equal(t, "test author", expected["author"])
	assert.Equal(t, "test title", expected["title"])
	assert.Equal(t, "some random text here", expected["text"])
	assert.NotEmpty(t, expected["uuid"])
}

func TestApiArticleUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// arrange
	ut := newArticleAPIUnitTest()
	articleUUID := uuid.New()

	assert.NoError(t, ut.repository.Create(context.Background(), &data.Article{
		UUID:      articleUUID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Title:     "test title",
		Author:    "test author",
		Text:      "some random text here",
	}))

	newArtile := map[string]interface{}{
		"author": "another author",
		"title":  "test title",
		"text":   "some random text here",
	}
	body, err := json.Marshal(newArtile)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/v1/articles/%s", articleUUID.String()), bytes.NewReader(body))
	req.Header.Add("Content-Type", "application/json")

	r := internal.RouterInitializer(ut.handler)

	// act
	r.ServeHTTP(rec, req)

	// assert
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

	expected := make(map[string]interface{})
	assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &expected))

	assert.Equal(t, "another author", expected["author"])
	assert.Equal(t, "test title", expected["title"])
	assert.Equal(t, "some random text here", expected["text"])
}

func TestApiArticleRead(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// arrange
	ut := newArticleAPIUnitTest()
	articleUUID := uuid.New()

	assert.NoError(t, ut.repository.Create(context.Background(), &data.Article{
		UUID:      articleUUID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Title:     "test title",
		Author:    "test author",
		Text:      "some random text here",
	}))

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/v1/articles/%s", articleUUID.String()), nil)
	req.Header.Add("Content-Type", "application/json")

	r := internal.RouterInitializer(ut.handler)

	// act
	r.ServeHTTP(rec, req)

	// assert
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

	expected := make(map[string]interface{})
	assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &expected))

	assert.Equal(t, "test author", expected["author"])
	assert.Equal(t, "test title", expected["title"])
	assert.Equal(t, "some random text here", expected["text"])
}

func TestApiArticleDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// arrange
	ut := newArticleAPIUnitTest()
	articleUUID := uuid.New()

	assert.NoError(t, ut.repository.Create(context.Background(), &data.Article{
		UUID:      articleUUID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Title:     "test title",
		Author:    "test author",
		Text:      "some random text here",
	}))

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/v1/articles/%s", articleUUID.String()), nil)
	req.Header.Add("Content-Type", "application/json")

	r := internal.RouterInitializer(ut.handler)

	// act
	r.ServeHTTP(rec, req)

	// assert
	assert.Equal(t, http.StatusNoContent, rec.Code)
}
