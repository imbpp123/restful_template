package presenter

import (
	"context"
	"net/http"
)

type (
	ApiVersion    string
	ApiVersionCtx struct{}
)

const (
	APIVersion1 ApiVersion = "v1"
	APIVersion2 ApiVersion = "v2"

	DefaultAPIVersion ApiVersion = APIVersion2
)

func ApiVersionMiddleware(version string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), ApiVersionCtx{}, ApiVersion(version)))

			next.ServeHTTP(w, r)
		})
	}
}

func getAPIVersion(r *http.Request) ApiVersion {
	apiVersion, ok := r.Context().Value(ApiVersionCtx{}).(ApiVersion)
	if !ok {
		return DefaultAPIVersion
	}

	return apiVersion
}
