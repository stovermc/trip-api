package app

import (
	"context"

	"github.com/stovermc/river-right-api/internal/trips/domain"
)

type Dependencies struct {
	GenerateID
	TripRepository
}

type GenerateID func() (string, error)

type TripRepository interface {
	Add(context.Context, domain.Trip) error
	Get(context.Context, string) (domain.Trip, error)
	Save(context.Context, domain.Trip) error
}


type Logger interface {
	Info(args ...interface{})
	Error(err error, args ...interface{})
}