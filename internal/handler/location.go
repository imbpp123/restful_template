package handler

import (
	"app/internal/data"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

const (
	paramLocationRiderID string = "riderID"
	paramMaxLocation     string = "maxlocation"
)

type (
	locationCreateRequest struct {
		Latitude  float64 `json:"lat" validate:"required"`
		Longitude float64 `json:"long" validate:"required"`
	}

	locationResponse struct {
		Latitude  float64 `json:"lat"`
		Longitude float64 `json:"long"`
	}

	listResponse struct {
		RiderID string             `json:"rider_id"`
		History []locationResponse `json:"history"`
	}

	locationService interface {
		Create(ctx context.Context, createLocation *data.CreateLocation) (*data.Location, error)
		List(ctx context.Context, params *data.ListLocation) ([]data.Location, error)
	}

	LocationHandler struct {
		locationService locationService
		validator       *validator.Validate
	}
)

func NewLocationHandler(
	locationService locationService,
	validator *validator.Validate,
) *LocationHandler {
	return &LocationHandler{
		locationService: locationService,
		validator:       validator,
	}
}

func LocationRouter(locationHandler *LocationHandler) http.Handler {
	r := chi.NewRouter()

	r.Route(fmt.Sprintf("/{%s}", paramLocationRiderID), func(r chi.Router) {
		r.Post("/now", locationHandler.Create())
		r.Get("/", locationHandler.List())
	})

	return r
}

func (h *LocationHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		riderID, err := h.getRiderUUID(r, paramLocationRiderID)
		if err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		request := &locationCreateRequest{}
		if err := json.NewDecoder(r.Body).Decode(request); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		if err := h.validator.Struct(request); err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		createData := &data.CreateLocation{
			RiderID:   riderID,
			Latitude:  request.Latitude,
			Longitude: request.Longitude,
		}
		_, err = h.locationService.Create(r.Context(), createData)
		if err != nil {
			render.Render(w, r, ErrByError(err))
			return
		}

		render.Status(r, http.StatusCreated)
	}
}

func (h *LocationHandler) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		riderID, err := h.getRiderUUID(r, paramLocationRiderID)
		if err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}

		var locationList []data.Location

		max := h.getRequestParamInteger(r, paramMaxLocation, data.LocationListAll)
		locationList, err = h.locationService.List(r.Context(), &data.ListLocation{
			RiderID:  riderID,
			MaxCount: max,
		})
		if err != nil {
			render.Render(w, r, ErrByError(err))
			return
		}

		response := listResponse{
			RiderID: riderID,
		}
		for _, item := range locationList {
			response.History = append(response.History, locationResponse{
				Latitude:  item.Latitude,
				Longitude: item.Longitude,
			})
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func (h *LocationHandler) getRiderUUID(r *http.Request, paramName string) (string, error) {
	idStr := chi.URLParam(r, paramName)
	if idStr == "" {
		return "", data.ErrParameterNotFound
	}

	return idStr, nil
}

func (h *LocationHandler) getRequestParamInteger(r *http.Request, paramName string, defaultValue int) int {
	numberStr := r.URL.Query().Get(paramName)
	if numberStr == "" {
		return defaultValue
	}

	number, err := strconv.Atoi(numberStr)
	if err != nil {
		return defaultValue
	}

	return number
}
