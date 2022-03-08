package stub

import "github.com/stovermc/river-right-api/internal/trips/app"

type stubLogger struct{}

func NewStubLogger() app.Logger {
	return &stubLogger{}
}

func (l *stubLogger) Info(args ...interface{})             {}
func (l *stubLogger) Error(err error, args ...interface{}) {}
