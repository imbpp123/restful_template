package presenter

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"

	"app/internal/data"
	v1 "app/internal/handler/presenter/v1"
	v2 "app/internal/handler/presenter/v2"
)

type (
	VersionRequest[T any] interface {
		Bind(r *http.Request) error
		TransformTo() (*T, error)
	}

	ArticleCreateRequest struct {
	}

	ArticleUpdateRequest struct {
	}

	ArticleListRequest struct {
	}

	ArticleResponse struct {
	}
)

var validate = validator.New()

func NewArticleCreateRequest() *ArticleCreateRequest {
	return &ArticleCreateRequest{}
}

func (a *ArticleCreateRequest) Bind(r *http.Request) (*data.CreateArticle, error) {
	versionMap := map[ApiVersion]VersionRequest[data.CreateArticle]{
		APIVersion1: &v1.ArticleCreateRequest{},
		APIVersion2: &v2.ArticleCreateRequest{},
	}

	return bindRequest[data.CreateArticle](r, versionMap)
}

func NewArticleUpdateRequest() *ArticleUpdateRequest {
	return &ArticleUpdateRequest{}
}

func (a *ArticleUpdateRequest) Bind(r *http.Request) (*data.UpdateArticle, error) {
	versionMap := map[ApiVersion]VersionRequest[data.UpdateArticle]{
		APIVersion1: &v1.ArticleUpdateRequest{},
		APIVersion2: &v2.ArticleUpdateRequest{},
	}

	return bindRequest[data.UpdateArticle](r, versionMap)
}

func NewArticleListRequest() *ArticleListRequest {
	return &ArticleListRequest{}
}

func (a *ArticleListRequest) Bind(r *http.Request) (*data.ArticleListParameters, error) {
	versionMap := map[ApiVersion]VersionRequest[data.ArticleListParameters]{
		APIVersion1: &v1.ArticleListRequest{},
		APIVersion2: &v2.ArticleListRequest{},
	}

	return bindRequest[data.ArticleListParameters](r, versionMap)
}

func NewArticleResponse() *ArticleResponse {
	return &ArticleResponse{}
}

func (a *ArticleResponse) Render(w http.ResponseWriter, r *http.Request, data *data.Article) {
	var payload render.Renderer

	switch getAPIVersion(r) {
	case APIVersion1:
		payload = v1.NewArticleResponse(data)
	case APIVersion2:
		payload = v2.NewArticleResponse(data)
	}

	render.Render(w, r, payload)
}

func bindRequest[T any](r *http.Request, versionMap map[ApiVersion]VersionRequest[T]) (*T, error) {
	apiVersion := getAPIVersion(r)

	request, ok := versionMap[apiVersion]
	if !ok {
		return nil, data.ErrUnsupportedAPIVersion
	}

	bindable, ok := any(request).(interface {
		Bind(*http.Request) error
	})
	if !ok {
		return nil, data.ErrUnsupportedBinder
	}

	if err := render.Bind(r, bindable); err != nil {
		return nil, fmt.Errorf("presenter.bindRequest: %w", err)
	}

	if err := validate.Struct(request); err != nil {
		return nil, fmt.Errorf("presenter.bindRequest: %w", err)
	}

	transformable, ok := any(request).(interface {
		TransformTo() (*T, error)
	})
	if !ok {
		return nil, data.ErrUnsupportedTransform
	}

	result, err := transformable.TransformTo()
	if err != nil {
		return nil, fmt.Errorf("presenter.bindRequest: %w", err)
	}

	return result, nil
}
