package mongo_test

import (
	"context"
	"testing"

	"github.com/stovermc/river-right-api/internal/trips/app"
	"github.com/stovermc/river-right-api/internal/trips/domain"
	"github.com/stovermc/river-right-api/internal/trips/frameworks/persistence/mongo"
	"github.com/stovermc/river-right-api/test"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	driver "go.mongodb.org/mongo-driver/mongo"
)

const tripsCollection = "trips"
func Test_TripRepository_Add_WhenTripDoesNotExist(t *testing.T) {
	client, db, tearDown := setUp(t)
	defer tearDown()

	tripID := test.NewRandomID()
	trip := domain.NewTrip(tripID, "test-trip-name")
	repo := mongo.NewTripRepository(client, db)
	err := repo.Add(context.Background(), trip)

	assert.Nil(t, err, "failed while saving trip")

	tripDoc := mongo.Trip{}
	err = client.Database(db).Collection(tripsCollection).FindOne(context.Background(), byIdFilter(tripID)).Decode(&tripDoc)
	assert.Nil(t, err, "failed while retrieving saved trip")

	assertModelEqualsDoc(t, trip, tripDoc)

	assertCountIs(t, client, db, tripID, 1)
}

func Test_TripRepository_Add_WhenTripAlreadyExists(t *testing.T) {
	client, db, tearDown := setUp(t)
	defer tearDown()

	want := app.ErrTripAlreadyExists

	tripID := test.NewRandomID()
	tripDoc := mongo.Trip{ID: tripID, Name: "test-trip-name"}

	_, err := client.Database(db).Collection(tripsCollection).InsertOne(context.Background(), tripDoc)
	assert.Nil(t, err, "failed inserting tripDoc")

	trip := domain.NewTrip(tripID, "test-trip-name")

	repo := mongo.NewTripRepository(client, db)
	err = repo.Add(context.Background(), trip)

	assert.ErrorIs(t, want, err)
}

func Test_TripRepository_Get_WhenTripAlreadyExists(t *testing.T) {
	client, db, tearDown := setUp(t)
	defer tearDown()

	tripID := test.NewRandomID()
	tripDoc := mongo.Trip{ID: tripID, Name: "test-trip-name"}

	_, err := client.Database(db).Collection(tripsCollection).InsertOne(context.Background(), tripDoc)
	assert.Nil(t, err, "failed inserting tripDoc")

	repo := mongo.NewTripRepository(client, db)
	got, err := repo.Get(context.Background(), tripID)
	assert.Nil(t, err, "failed getting trip")

	assertModelEqualsDoc(t, got, tripDoc)
}

func Test_TripRepository_Get_WhenTripDoesNotExist(t *testing.T) {
	client, db, tearDown := setUp(t)
	defer tearDown()

	want := app.ErrTripNotFound

	tripID := test.NewRandomID()
	repo := mongo.NewTripRepository(client, db)
	_, err := repo.Get(context.Background(), tripID)

	assert.ErrorIs(t, want, err)
}

func Test_TripRepository_Save_WhenTripAlreadyExists(t *testing.T) {
	client, db, tearDown := setUp(t)
	defer tearDown()

	tripID := test.NewRandomID()
	tripDoc := mongo.Trip{ID: tripID, Name: "test-trip-name"}

	_, err := client.Database(db).Collection(tripsCollection).InsertOne(context.Background(), tripDoc)
	assert.Nil(t, err, "failed inserting tripDoc")

	repo := mongo.NewTripRepository(client, db)
	trip := domain.NewTrip(tripID, "test-updated-trip-name")

	err = repo.Save(context.Background(), trip)
	assert.Nil(t, err, "failed while saving trip")

	updatedTripDoc := mongo.Trip{}
	err = client.Database(db).Collection(tripsCollection).FindOne(context.Background(), byIdFilter(tripID)).Decode(&updatedTripDoc)
	assert.Nil(t, err, "failed while retrieving saved trip")

	assertModelEqualsDoc(t, trip, updatedTripDoc)

	assertCountIs(t, client, db, tripID, 1)
}

func Test_TripRepository_Save_WhenTripDoesNotExists(t *testing.T) {
	client, db, tearDown := setUp(t)
	defer tearDown()

	want := app.ErrTripNotFound

	tripID := test.NewRandomID()
	trip := domain.NewTrip(tripID, "test-trip-name")

	repo := mongo.NewTripRepository(client, db)
	err := repo.Save(context.Background(), trip)

	assert.ErrorIs(t, want, err)
}

// Test Helper functions
func byIdFilter(id string) bson.D {
	return bson.D{{"identifier", id}}
}

func assertCountIs(t *testing.T, client *driver.Client, db string, tripID string, want int64) {
	t.Helper()
	count, err := client.Database(db).Collection(tripsCollection).CountDocuments(context.Background(), byIdFilter(tripID))
	assert.Nil(t, err, "failed counting docs")
	assert.Equal(t, count, want, "doc count")
}

func assertModelEqualsDoc(t *testing.T, m domain.Trip, d mongo.Trip) {
	t.Helper()
	assert.Equal(t, m.ID(), d.ID, "trip.ID()")
	assert.Equal(t, m.Name(), d.Name, "trip.Name()")
}
