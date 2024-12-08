package data

import (
	"time"

	"github.com/google/uuid"
)

type (
	Article struct {
		UUID      uuid.UUID
		CreatedAt time.Time
		UpdatedAt time.Time
		Title     string
		Author    string
		Text      string
	}

	CreateArticle struct {
		UUID   uuid.UUID
		Title  string
		Author string
		Text   string
	}

	UpdateArticle struct {
		Title  string
		Author string
		Text   string
	}

	ArticleListParameters struct {
		Pagination

		Author *string
	}
)

func NewArticle(uuid uuid.UUID) *Article {
	return &Article{
		UUID: uuid,
	}
}
