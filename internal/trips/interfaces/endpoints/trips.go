package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/stovermc/river-right-api/internal/trips/app"
	"github.com/stovermc/river-right-api/internal/trips/domain"
	"github.com/stovermc/river-right-api/internal/trips/interfaces/presenters"
)

func MakeAddTrip(d *app.Dependencies) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(presenters.AddTripRequest)
		tripID := d.GenerateID()

		trip := domain.NewTrip(tripID, req.Name)

		d.TripRepository.Add(ctx, trip)

		return presenters.AddTripResponse{
			Name: req.Name,
		}, nil
	}
}
