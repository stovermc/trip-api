package inmemory

import (
	"context"
	"sync"

	"github.com/stovermc/river-right-api/internal/trips/app"
	"github.com/stovermc/river-right-api/internal/trips/domain"
)

type TripRepo struct {
	Trips map[string]domain.Trip
	sync.Mutex
}

func NewTripRepository() app.TripRepository {
	m := make(map[string]domain.Trip)

	return &TripRepo{
		Trips: m,
	}
}

func (tr *TripRepo) Add(ctx context.Context, t domain.Trip) error {
	tr.Lock()
	if _, ok := tr.Trips[t.ID()]; ok {
		tr.Unlock()
		return app.ErrTripAlreadyExists
	}

	tr.Trips[t.ID()] = t
	tr.Unlock()

	return nil
}

func (tr *TripRepo) Get(ctx context.Context, id string) (domain.Trip, error) {
	tr.Lock()
	t, ok := tr.Trips[id]
	if !ok {
		tr.Unlock()
		return nil, app.ErrTripNotFound
	}
	tr.Unlock()

	return t, nil
}

func (tr *TripRepo) Save(ctx context.Context, t domain.Trip) error {
	tr.Lock()
	if _, ok := tr.Trips[t.ID()]; !ok {
		tr.Unlock()
		return app.ErrTripNotFound
	}

	tr.Trips[t.ID()] = t
	tr.Unlock()

	return nil
}