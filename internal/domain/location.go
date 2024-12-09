package domain

import (
	"app/internal/data"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type (
	repository interface {
		Create(ctx context.Context, location *data.Location) error
		FindByRiderUUID(ctx context.Context, riderUUID string, maxCount int) ([]data.Location, error)
	}

	LocationService struct {
		repository repository
	}
)

func NewLocationService(repository repository) *LocationService {
	return &LocationService{
		repository: repository,
	}
}

func (l *LocationService) Create(ctx context.Context, createLocation *data.CreateLocation) (*data.Location, error) {
	newLocation := &data.Location{
		UUID:      uuid.New(),
		CreatedAt: time.Now(),
		RiderID:   createLocation.RiderID,
		Latitude:  createLocation.Latitude,
		Longitude: createLocation.Longitude,
	}

	if err := l.repository.Create(ctx, newLocation); err != nil {
		return nil, fmt.Errorf("domain.LocationService.Create: %w", err)
	}

	return newLocation, nil
}

func (l *LocationService) List(ctx context.Context, params *data.ListLocation) ([]data.Location, error) {
	result, err := l.repository.FindByRiderUUID(ctx, params.RiderID, params.MaxCount)
	if err != nil {
		return nil, fmt.Errorf("domain.LocationService.List: %w", err)
	}

	return result, nil
}
