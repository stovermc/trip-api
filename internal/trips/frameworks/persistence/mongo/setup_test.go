package mongo_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stovermc/river-right-api/internal/trips/frameworks/persistence/mongo"
	"github.com/stovermc/river-right-api/test"
	"github.com/stovermc/river-right-api/test/stub"
	"github.com/stretchr/testify/assert"
	driver "go.mongodb.org/mongo-driver/mongo"
)

type tearDown func()

const dbPrefix = "trips_test_"

func setUp(t *testing.T) (client *driver.Client, db string, td tearDown) {
	t.Helper()

	db = dbPrefix + test.NewRandomID()
	logger := stub.NewStubLogger()
	client, disconnect, err := mongo.NewClient(logger, mongo.Connection{
		Host: "localhost",
		Port: "27017",
		Datbase: "admin",
		Username: "root",
		Password: "SuperSecret123",
	})

	assert.Nil(t, err, "failed to create new mongo client")
	
	err = client.Database(db).Drop(context.Background())
	assert.Nil(t, err, fmt.Sprintf("failed to set up db: %s", db))

	td = func() {
		err = client.Database(db).Drop(context.Background())
		assert.Nil(t, err, fmt.Sprintf("failed to tear down db: %s", db))
		disconnect()
	}

	return
}