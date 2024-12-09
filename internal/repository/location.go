package repository

import (
	"app/internal/data"
	"context"
	"sync"
)

type Location struct {
	data map[string][]data.Location

	mu sync.RWMutex
}

func NewLocation() *Location {
	return &Location{
		data: make(map[string][]data.Location),
	}
}

func (l *Location) Create(ctx context.Context, location *data.Location) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.data[location.RiderID] = append(l.data[location.RiderID], *location)

	return nil
}

func (l *Location) FindByRiderUUID(ctx context.Context, riderUUID string, maxCount int) ([]data.Location, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	locations, ok := l.data[riderUUID]
	if !ok {
		return nil, data.ErrLocationNotFound
	}

	n := len(locations)
	if n < maxCount {
		maxCount = n
	}

	result := make([]data.Location, 0)
	j := 0
	for i := n - 1; i >= maxCount; i-- {
		result = append(result, locations[i])
		j++
	}

	return result, nil
}
