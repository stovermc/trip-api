package inmemory_test

import (
	"context"
	"testing"

	"github.com/stovermc/river-right-api/internal/trips/app"
	"github.com/stovermc/river-right-api/internal/trips/domain"
	inmemory "github.com/stovermc/river-right-api/internal/trips/frameworks/persistence/inMemory"
	"github.com/stovermc/river-right-api/test"
	"github.com/stretchr/testify/assert"
)


func Test_TripRepository_Add_WhenTripDoesNotExist(t *testing.T) {
	assert :=assert.New(t)

	id := test.NewRandomID()
	name := test.NewRandomID()
	trip := domain.NewTrip(id, name)

	tripRepo := inmemory.NewTripRepository()
	err := tripRepo.Add(context.Background(), trip)
	assert.Nil(err, "failed while saving trip")

	inmemRepo, ok := tripRepo.(*inmemory.TripRepo)
	assert.True(ok, "failed while type casting to inmemory.TripRepo")

	persistedTrip, ok := inmemRepo.Trips[id]
  assert.True(ok, "trip not persisted in repo")
	assert.Equal(trip, persistedTrip, "trip not persisted in repo")
}

func Test_TripRepository_Add_WhenTripAlreadyExists(t *testing.T) {
	assert :=assert.New(t)
	want := app.ErrTripAlreadyExists

	id := test.NewRandomID()
	existingTrip := domain.NewTrip(id, "test-name")
	trip := domain.NewTrip(id, "test-some-other-name")

	tripRepo := inmemory.NewTripRepository()

	err := tripRepo.Add(context.Background(), existingTrip)
	assert.Nil(err, "failed while saving trip")

	err = tripRepo.Add(context.Background(), trip)
	assert.Equal(err, want)
}

func Test_TripRepository_Get_WhenTripAlreadyExists(t *testing.T) {
	assert :=assert.New(t)

	id := test.NewRandomID()
	name := test.NewRandomID()
	trip := domain.NewTrip(id, name)
	tripRepo := inmemory.NewTripRepository()
	err := tripRepo.Add(context.Background(), trip)
	assert.Nil(err, "failed while saving trip")

	persistedTrip, err := tripRepo.Get(context.Background(), id)

	assert.Equal(persistedTrip, trip)
}

func Test_TripRepository_Get_WhenTripDoesNotExist(t *testing.T) {
	assert :=assert.New(t)
	want := app.ErrTripNotFound

	id := test.NewRandomID()
	tripRepo := inmemory.NewTripRepository()

	_, err := tripRepo.Get(context.Background(), id)

	assert.Equal(err, want)
}

func Test_TripRepository_Save_WhenTripAlreadyExists(t *testing.T) {
	assert :=assert.New(t)

	id := test.NewRandomID()
	name := test.NewRandomID()
	trip := domain.NewTrip(id, name)
	tripRepo := inmemory.NewTripRepository()
	err := tripRepo.Add(context.Background(), trip)
	assert.Nil(err, "failed while saving trip")

	updatedTrip := domain.NewTrip(id, "updated-name")
	err = tripRepo.Save(context.Background(), updatedTrip)
	assert.Nil(err, "failed while saving trip")

	inmemRepo, ok := tripRepo.(*inmemory.TripRepo)
	assert.True(ok, "failed while type casting to inmemory.TripRepo")

	persistedTrip, ok := inmemRepo.Trips[id]
  assert.True(ok, "trip not persisted in repo")
	assert.Equal(updatedTrip, persistedTrip, "trip not persisted in repo")
}

func Test_TripRepository_Save_WhenTripDoesNotExists(t *testing.T) {
	assert :=assert.New(t)
	want := app.ErrTripNotFound

	id := test.NewRandomID()
	name := test.NewRandomID()
	trip := domain.NewTrip(id, name)
	tripRepo := inmemory.NewTripRepository()

	err := tripRepo.Save(context.Background(), trip)

	assert.Equal(err, want)
}