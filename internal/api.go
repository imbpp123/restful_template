package internal

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"app/internal/domain"
	"app/internal/handler"
	"app/internal/handler/presenter"
	"app/internal/repository"
)

func RouterInitializer() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.Heartbeat("/ping"))

	/*
		articlePresenter := presenter.NewArticle()
		errorPresenter := presenter.NewError()

		articleRepository := repository.NewArticle()
		articleService := domain.NewArticle(
			articleRepository,
		)
		articleHandler := handler.NewArticleHandler(
			articleService,
			articlePresenter,
			errorPresenter,
		)
		articleRouter := handler.ArticleRouter(articleHandler)

		// API version 2.
		r.Route("/v2", func(r chi.Router) {
			r.Use(presenter.ApiVersion(presenter.APIVersion2))
			r.Mount("/articles", articleRouter)
		})

		// API version 1.
		r.Route("/v1", func(r chi.Router) {
			r.Use(presenter.ApiVersion(presenter.APIVersion1))
			r.Mount("/articles", articleRouter)
		})
	*/

	return r
}

func InitializeArticleHandler() *handler.ArticleHandler {
	errorPresenter := presenter.NewError()
	articlePresenter := presenter.NewArticleResponse()

	articleRepository := repository.NewArticle()
	articleService := domain.NewArticle(
		articleRepository,
	)

	return handler.NewArticleHandler(
		articleService,
		articlePresenter,
		errorPresenter,
	)
}
