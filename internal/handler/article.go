package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"app/internal/data"
	"app/internal/handler/presenter"
)

const paramArticleID string = "articleID"

type (
	articlePresenter interface {
		Render(w http.ResponseWriter, r *http.Request, data *data.Article)
	}

	errorPresenter interface {
		Render(w http.ResponseWriter, r *http.Request, err error)
	}

	articleService interface {
		Create(ctx context.Context, object *data.CreateArticle) (uuid.UUID, error)
		List(ctx context.Context, parameters *data.ArticleListParameters) ([]data.Article, error)
		LoadByID(ctx context.Context, uuid uuid.UUID) (*data.Article, error)
		Update(ctx context.Context, uuid uuid.UUID, object *data.UpdateArticle) error
		DeleteByUUID(ctx context.Context, uuid uuid.UUID) error
	}

	ArticleHandler struct {
		articleService   articleService
		articlePresenter articlePresenter
		errorPresenter   errorPresenter
	}
)

func ArticleRouter(articleHandler *ArticleHandler) http.Handler {
	r := chi.NewRouter()

	//r.Method("GET", "/", articleHandler.List())
	r.Method("POST", "/", articleHandler.Create())
	r.Route(fmt.Sprintf("/{%s}", paramArticleID), func(r chi.Router) {
		r.Get("/", articleHandler.View())
		r.Put("/", articleHandler.Update())
		r.Delete("/", articleHandler.Delete())
	})

	return r
}

func NewArticleHandler(
	articleService articleService,
	articlePresenter articlePresenter,
	errorPresenter errorPresenter,
) *ArticleHandler {
	return &ArticleHandler{
		articleService:   articleService,
		articlePresenter: articlePresenter,
		errorPresenter:   errorPresenter,
	}
}

func (h *ArticleHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		binder := presenter.NewArticleCreateRequest()

		articleRequest, err := binder.Bind(r)
		if err != nil {
			h.errorPresenter.Render(w, r, fmt.Errorf("Handler.ArticleHandler.Create: %w", err))
			return
		}

		articleUUID, err := h.articleService.Create(r.Context(), articleRequest)
		if err != nil {
			h.errorPresenter.Render(w, r, fmt.Errorf("Handler.ArticleHandler.View: %w", err))
			return
		}

		article, err := h.articleService.LoadByID(r.Context(), articleUUID)
		if err != nil {
			h.errorPresenter.Render(w, r, fmt.Errorf("Handler.ArticleHandler.View: %w", err))
			return
		}

		h.articlePresenter.Render(w, r, article)
	}
}

func (h *ArticleHandler) View() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		articleUUID, err := h.getUUID(r)
		if err != nil {
			h.errorPresenter.Render(w, r, fmt.Errorf("ArticleHandler.View: %w", err))
			return
		}

		article, err := h.articleService.LoadByID(r.Context(), articleUUID)
		if err != nil {
			h.errorPresenter.Render(w, r, fmt.Errorf("ArticleHandler.View: %w", err))
			return
		}

		h.articlePresenter.Render(w, r, article)
	}
}

func (h *ArticleHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		binder := presenter.NewArticleUpdateRequest()

		articleUpdate, err := binder.Bind(r)
		if err != nil {
			h.errorPresenter.Render(w, r, fmt.Errorf("ArticleHandler.View: %w", err))
			return
		}

		if err := h.articleService.Update(r.Context(), uuid.UUID{}, articleUpdate); err != nil {
			h.errorPresenter.Render(w, r, fmt.Errorf("ArticleHandler.View: %w", err))
			return
		}

		article, err := h.articleService.LoadByID(r.Context(), articleUpdate.UUID)
		if err != nil {
			h.errorPresenter.Render(w, r, fmt.Errorf("ArticleHandler.View: %w", err))
			return
		}

		h.articlePresenter.Render(w, r, article)
	}
}

func (h *ArticleHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid, err := h.getUUID(r)
		if err != nil {
			h.errorPresenter.Render(w, r, fmt.Errorf("ArticleHandler.View: %w", err))
			return
		}

		if err := h.articleService.DeleteByUUID(r.Context(), uuid); err != nil {
			h.errorPresenter.Render(w, r, fmt.Errorf("ArticleHandler.View: %w", err))
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func (h *ArticleHandler) List() func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		return nil
	}
}

func (h *ArticleHandler) getUUID(r *http.Request) (uuid.UUID, error) {
	id := chi.URLParam(r, paramArticleID)
	if id == "" {
		return uuid.UUID{}, data.ErrParameterNotFound
	}

	articleUUID, err := uuid.Parse(id)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("Handler.ArticleUUID: %w", err)
	}

	return articleUUID, nil
}
