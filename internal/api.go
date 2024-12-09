package internal

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"

	"app/internal/domain"
	"app/internal/handler"
	"app/internal/repository"
)

func RouterInitializer(articleHandler *handler.ArticleHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.Heartbeat("/ping"))

	articleRouter := handler.ArticleRouter(articleHandler)

	// API version 1.
	r.Route("/v1", func(r chi.Router) {
		r.Mount("/articles", articleRouter)
	})

	locationRepository := repository.NewLocation()
	locationService := domain.NewLocationService(locationRepository)
	locationHandler := handler.NewLocationHandler(
		locationService,
		validator.New(),
	)

	r.Mount("/location", handler.LocationRouter(locationHandler))

	return r
}

func InitializeArticleHandler() *handler.ArticleHandler {
	articleRepository := repository.NewArticle()
	articleService := domain.NewArticle(
		articleRepository,
	)

	return handler.NewArticleHandler(
		articleService,
		validator.New(),
	)
}
