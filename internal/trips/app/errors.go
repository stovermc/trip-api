package app

import "errors"

var (
	ErrTripNotFound = errors.New("trip not found")
	ErrTripAlreadyExists = errors.New("trip id already exists")
)