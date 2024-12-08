package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type validatedRequest struct{}

var validate = validator.New()

func ValidateRequest[T any](next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req T

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, fmt.Sprintf("invalid JSON: %v", err), http.StatusBadRequest)
			return
		}

		if err := validate.Struct(req); err != nil {
			http.Error(w, fmt.Sprintf("validation error: %v", err), http.StatusUnprocessableEntity)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), validatedRequest{}, req)))
	})
}
