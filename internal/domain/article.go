package domain

import (
	"context"
	"errors"
	"fmt"
	"time"

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
	return nil, errors.New("not implemented")
}

func (a *Article) LoadByID(ctx context.Context, uuid uuid.UUID) (*data.Article, error) {
	article, err := a.repository.FindByUUID(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("Domain.Article.LoadByID: %w", err)
	}

	return article, nil
}

func (a *Article) Create(ctx context.Context, createArticle *data.CreateArticle) (*data.Article, error) {
	newArticle := &data.Article{
		UUID:      createArticle.UUID,
		CreatedAt: time.Now(),
		Title:     createArticle.Title,
		Author:    createArticle.Author,
		Text:      createArticle.Text,
	}

	if err := a.repository.Create(ctx, newArticle); err != nil {
		return nil, fmt.Errorf("Domain.Article.Save: %w", err)
	}

	return newArticle, nil
}

func (a *Article) Update(ctx context.Context, uuid uuid.UUID, updateArticle *data.UpdateArticle) (*data.Article, error) {
	article, err := a.repository.FindByUUID(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("Domain.Article.LoadByID: %w", err)
	}

	article.Author = updateArticle.Author
	article.Text = updateArticle.Text
	article.Title = updateArticle.Title
	article.UpdatedAt = time.Now()

	if err := a.repository.Update(ctx, article); err != nil {
		return nil, fmt.Errorf("Domain.Article.Save: %w", err)
	}

	return article, nil
}

func (a *Article) DeleteByUUID(ctx context.Context, uuid uuid.UUID) error {
	if err := a.repository.DeleteByUUID(ctx, uuid); err != nil {
		return fmt.Errorf("Domain.Article.DeleteByID: %w", err)
	}

	return nil
}
