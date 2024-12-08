package data

import (
	"time"

	"github.com/google/uuid"
)

type (
	Article struct {
		ID        uuid.UUID
		CreatedAt time.Time
		UpdatedAt time.Time
		Title     string
		Author    string
		Text      string
		Tags      []string
	}

	CreateArticle struct {
		UUID   uuid.UUID
		Title  string
		Author string
		Text   string
		Tags   []string
	}

	UpdateArticle struct {
		UUID   uuid.UUID
		Title  string
		Author string
		Text   string
		Tags   []string
	}

	ArticleListParameters struct {
		Pagination

		Author *string
	}
)

func NewArticle(uuid uuid.UUID) *Article {
	return &Article{
		ID: uuid,
	}
}
