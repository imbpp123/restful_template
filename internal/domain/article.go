package domain

import (
	"context"

	"github.com/google/uuid"

	"app/internal/data"
)

type (
	ArticleRepository interface {
		FindByUUID(ctx context.Context, uuid uuid.UUID) (*data.Article, error)
		DeleteByUUID(ctx context.Context, uuid uuid.UUID) error
		List(ctx context.Context, params *data.ArticleListParameters) ([]data.Article, error)
		Create(ctx context.Context, article *data.Article) error
		Update(ctx context.Context, article *data.Article) error
	}

	Article struct {
		repository ArticleRepository
	}
)

func NewArticle(repository ArticleRepository) *Article {
	return &Article{
		repository: repository,
	}
}

func (a *Article) List(ctx context.Context, parameters *data.ArticleListParameters) ([]data.Article, error) {
	return nil, nil
}

func (a *Article) LoadByID(ctx context.Context, uuid uuid.UUID) (*data.Article, error) {
	/*
		data, err := a.repository.FindByUUID(ctx, id)
		if err != nil {
			return nil, fmt.Errorf("Domain.Article.LoadByID: %w", err)
		}
	*/

	return nil, nil
}

func (a *Article) Create(ctx context.Context, article *data.CreateArticle) (uuid.UUID, error) {
	/*
		article.CreatedAt = time.Now()

		if err := a.repository.Create(ctx, article); err != nil {
			return fmt.Errorf("Domain.Article.Save: %w", err)
		}
	*/

	return uuid.UUID{}, nil
}

func (a *Article) Update(ctx context.Context, uuid uuid.UUID, article *data.UpdateArticle) error {
	/*
		article.UpdatedAt = time.Now()

		if err := a.repository.Update(ctx, article); err != nil {
			return fmt.Errorf("Domain.Article.Save: %w", err)
		}
	*/

	return nil
}

func (a *Article) DeleteByUUID(ctx context.Context, uuid uuid.UUID) error {
	/*
		if err := a.repository.DeleteByUUID(ctx, id); err != nil {
			return fmt.Errorf("Domain.Article.DeleteByID: %w", err)
		}
	*/

	return nil
}
