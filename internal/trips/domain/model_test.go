package domain_test

import (
	"testing"

	"github.com/stovermc/river-right-api/internal/trips/domain"
	"github.com/stovermc/river-right-api/test"
	"github.com/stretchr/testify/assert"
)

func Test_Trip_Attributes(t *testing.T) {
	assert.New(t)
	id := test.NewRandomID()
	name := test.NewRandomID()

	trip := domain.NewTrip(id, name)

	assert.Equal(t, trip.ID(), id)
	assert.Equal(t, trip.Name(), name)
}
