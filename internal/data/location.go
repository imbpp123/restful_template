package data

import (
	"time"

	"github.com/google/uuid"
)

const (
	LocationListAll int = 0
)

type (
	Location struct {
		UUID      uuid.UUID
		CreatedAt time.Time
		RiderID   string
		Latitude  float64
		Longitude float64
	}

	CreateLocation struct {
		RiderID   string
		Longitude float64
		Latitude  float64
	}

	ListLocation struct {
		RiderID  string
		MaxCount int
	}
)
