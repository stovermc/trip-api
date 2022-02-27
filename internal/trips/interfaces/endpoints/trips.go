package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/stovermc/river-right-api/internal/trips/app"
)

func MakeAddTrip(d *app.Dependencies) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// add presenter for request json -> domain AddTripRequest
		// tripID, err := d.GenerateID()
		// if err != nil {
		// 	return nil, fmt.Errorf("could not generate id for trip: %w", err)
		// }

		return nil, nil
	}
}

	// return func(ctx context.Context, request interface{}) (interface{}, error) {
	// 	req := request.(presenters.AddGifterRequest)

	// 	gifterID, err := d.GenerateID()
	// 	if err != nil {
	// 		return nil, fmt.Errorf("could not generate id: %w", err)
	// 	}

	// 	cmd := app.NewCommandMessage(domain.AddGifterCommand{
	// 		Name:     req.Name,
	// 		GifterID: gifterID,
	// 		GroupID:  req.GroupID,
	// 	})

	// 	err = d.MessageBus.Handle(ctx, cmd)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	return presenters.AddGifterResponse{
	// 		GifterID: gifterID,
	// 		GroupID:  req.GroupID,
	// 		Name:     req.Name,
	// 	}, nil
	// }