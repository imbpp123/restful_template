package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"app/internal/data"
)

const paramArticleID string = "articleID"

type (
	articleCreateRequest struct {
		Author string `json:"author" validate:"required"`
		Title  string `json:"title" validate:"required"`
		Text   string `json:"text" validate:"required"`
	}

	articleUpdateRequest struct {
		UUID   uuid.UUID `json:"uuid"  validate:"required"`
		Author string    `json:"author"  validate:"required"`
		Title  string    `json:"title"  validate:"required"`
		Text   string    `json:"text"  validate:"required"`
	}

	articleListRequest struct {
		Author string `json:"author"  validate:"required"`
		Title  string `json:"title"  validate:"required"`
		Text   string `json:"text"  validate:"required"`
	}

	articleResponse struct {
		UUID   string `json:"uuid"`
		Author string `json:"author"`
		Title  string `json:"title"`
		Text   string `json:"text"`
	}

	articleService interface {
		Create(ctx context.Context, object *data.CreateArticle) (*data.Article, error)
		List(ctx context.Context, parameters *data.ArticleListParameters) ([]data.Article, error)
		LoadByID(ctx context.Context, uuid uuid.UUID) (*data.Article, error)
		Update(ctx context.Context, uuid uuid.UUID, object *data.UpdateArticle) (*data.Article, error)
		DeleteByUUID(ctx context.Context, uuid uuid.UUID) error
	}

	ArticleHandler struct {
		articleService articleService
		validator      *validator.Validate
	}
)

func ArticleRouter(articleHandler *ArticleHandler) http.Handler {
	r := chi.NewRouter()

	r.Get("/", articleHandler.List())
	r.Post("/", articleHandler.Create())
	r.Route(fmt.Sprintf("/{%s}", paramArticleID), func(r chi.Router) {
		r.Get("/", articleHandler.View())
		r.Put("/", articleHandler.Update())
		r.Delete("/", articleHandler.Delete())
	})

	return r
}

func NewArticleHandler(
	articleService articleService,
	validator *validator.Validate,
) *ArticleHandler {
	return &ArticleHandler{
		articleService: articleService,
		validator:      validator,
	}
}

func (h *ArticleHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		createRequest := &articleCreateRequest{}

		if err := json.NewDecoder(r.Body).Decode(createRequest); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		if err := h.validator.Struct(createRequest); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		articleCreate := &data.CreateArticle{
			UUID:   uuid.New(),
			Author: createRequest.Author,
			Title:  createRequest.Title,
			Text:   createRequest.Text,
		}

		newArticle, err := h.articleService.Create(r.Context(), articleCreate)
		if err != nil {
			render.Render(w, r, ErrByError(err))
			return
		}

		w.WriteHeader(http.StatusCreated)
		render.Render(w, r, newArticleResponse(newArticle))
	}
}

func (h *ArticleHandler) View() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		articleUUID, err := h.getUUID(r, paramArticleID)
		if err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		article, err := h.articleService.LoadByID(r.Context(), articleUUID)
		if err != nil {
			render.Render(w, r, ErrByError(err))
			return
		}

		render.Render(w, r, newArticleResponse(article))
	}
}

func (h *ArticleHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		articleUUID, err := h.getUUID(r, paramArticleID)
		if err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		request := &articleUpdateRequest{
			UUID: articleUUID,
		}

		if err := json.NewDecoder(r.Body).Decode(request); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		if err := h.validator.Struct(request); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		updateArticle := &data.UpdateArticle{
			Author: request.Author,
			Title:  request.Title,
			Text:   request.Text,
		}

		updatedArticle, err := h.articleService.Update(r.Context(), articleUUID, updateArticle)
		if err != nil {
			render.Render(w, r, ErrByError(err))
			return
		}

		w.WriteHeader(http.StatusOK)
		render.Render(w, r, newArticleResponse(updatedArticle))
	}
}

func (h *ArticleHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid, err := h.getUUID(r, paramArticleID)
		if err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		if err := h.articleService.DeleteByUUID(r.Context(), uuid); err != nil {
			render.Render(w, r, ErrByError(err))
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func (h *ArticleHandler) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//articles, err := h.articleService.List(r.Context(), &data.ArticleListParameters{})

		//h.articleListPresenter.Render(w, r, articles)
	}
}

func (h *ArticleHandler) getUUID(r *http.Request, paramName string) (uuid.UUID, error) {
	id := chi.URLParam(r, paramName)
	if id == "" {
		return uuid.UUID{}, data.ErrParameterNotFound
	}

	articleUUID, err := uuid.Parse(id)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("Handler.getUUID: %w", err)
	}

	return articleUUID, nil
}

func newArticleResponse(article *data.Article) *articleResponse {
	return &articleResponse{
		UUID:   article.UUID.String(),
		Author: article.Author,
		Title:  article.Title,
		Text:   article.Text,
	}
}

func (ar *articleResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
