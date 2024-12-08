package v1

import (
	"net/http"

	"github.com/google/uuid"

	"app/internal/data"
)

type (
	ArticleCreateRequest struct {
		Author string `json:"author" validate:"required"`
		Title  string `json:"title" validate:"required"`
		Text   string `json:"text" validate:"required"`
	}

	ArticleUpdateRequest struct {
		UUID   uuid.UUID `json:"uuid"  validate:"required"`
		Author string    `json:"author"  validate:"required"`
		Title  string    `json:"title"  validate:"required"`
		Text   string    `json:"text"  validate:"required"`
	}

	ArticleListRequest struct {
		Author string `json:"author"  validate:"required"`
		Title  string `json:"title"  validate:"required"`
		Text   string `json:"text"  validate:"required"`
	}

	ArticleResponse struct {
		UUID     uuid.UUID `json:"uuid"`
		Author string    `json:"author"`
		Title  string    `json:"title"`
		Text   string    `json:"text"`
	}
)

func (a *ArticleCreateRequest) Bind(r *http.Request) error {
	return nil
}

func (a *ArticleCreateRequest) TransformTo() (*data.CreateArticle, error) {
	return &data.CreateArticle{
		Author: a.Author,
		Title:  a.Title,
		Text:   a.Text,
	}, nil
}

func (a *ArticleUpdateRequest) Bind(r *http.Request) error {
	return nil
}

func (a *ArticleUpdateRequest) TransformTo() (*data.UpdateArticle, error) {
	return &data.UpdateArticle{
		Author: a.Author,
		Title:  a.Title,
		Text:   a.Text,
	}, nil
}

func (a *ArticleListRequest) Bind(r *http.Request) error {
	return nil
}

func (a *ArticleListRequest) TransformTo() (*data.ArticleListParameters, error) {
	return &data.ArticleListParameters{}, nil
}

func NewArticleResponse(article *data.Article) *ArticleResponse {
	return &ArticleResponse{
		ID:     article.ID,
		Author: article.Author,
		Title:  article.Title,
		Text:   article.Text,
	}
}

func (a *ArticleResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
