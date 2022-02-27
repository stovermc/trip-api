package interfaces

import (
	"context"
	"encoding/json"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/stovermc/river-right-api/internal/trips/app"
	"github.com/stovermc/river-right-api/internal/trips/interfaces/endpoints"
)

func MakeHttpHandler(logger app.Logger, d *app.Dependencies) http.Handler {
	r := mux.NewRouter()

	r.Methods("POST").Path("/trips").Handler(kithttp.NewServer(
		endpoints.MakeAddTrip(d), 
		decodeAddTripRequest, 
		encodeAddTripResponse,
		// add error handling options...,
		))

	return r
}

func decodeAddTripRequest(c context.Context, r *http.Request) (interface{}, error) {
	// change to presenter
	var request interface{}

	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return nil, err
	}

	return request, nil
}

func encodeAddTripResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
