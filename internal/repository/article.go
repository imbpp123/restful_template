package repository

import (
	"context"
	"sync"

	"app/internal/data"

	"github.com/google/uuid"
)

type Article struct {
	data map[string]data.Article

	m sync.RWMutex
}

func NewArticle() *Article {
	return &Article{
		data: make(map[string]data.Article),
	}
}

func (a *Article) FindByUUID(ctx context.Context, id uuid.UUID) (*data.Article, error) {
	a.m.RLock()
	defer a.m.RUnlock()

	idStr := id.String()

	if article, ok := a.data[idStr]; ok {
		copy := article
		return &copy, nil
	}

	return nil, data.ErrArticleNotFound
}

func (a *Article) DeleteByUUID(ctx context.Context, id uuid.UUID) error {
	a.m.Lock()
	defer a.m.Unlock()

	delete(a.data, id.String())

	return nil
}

func (a *Article) List(ctx context.Context, params *data.ArticleListParameters) ([]data.Article, error) {
	return nil, nil
}

func (a *Article) Create(ctx context.Context, article *data.Article) error {
	a.m.Lock()
	defer a.m.Unlock()

	idStr := article.UUID.String()
	a.data[idStr] = *article

	return nil
}

func (a *Article) Update(ctx context.Context, article *data.Article) error {
	a.m.Lock()
	defer a.m.Unlock()

	idStr := article.UUID.String()

	var (
		dbArticle data.Article
		ok        bool
	)

	if dbArticle, ok = a.data[idStr]; !ok {
		return data.ErrArticleNotFound
	}

	dbArticle.Author = article.Author
	dbArticle.Title = article.Title
	dbArticle.Text = article.Text
	dbArticle.UpdatedAt = article.UpdatedAt

	a.data[idStr] = dbArticle

	return nil
}
