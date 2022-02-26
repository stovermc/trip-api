package log

import (
	"fmt"

	kitlog "github.com/go-kit/kit/log"
	"github.com/stovermc/river-right-api/internal/trips/app"
)

type logger struct {
	logger kitlog.Logger
}



func NewLogger(kl kitlog.Logger) app.Logger {
	return &logger{
		logger: kl,
	}
}

func (l *logger) Info(args ...interface{}) {
	l.logger.Log("msg", fmt.Sprint(args...))
}

func (l *logger) Error(err error, args ...interface{}) {
	l.logger.Log("msg", fmt.Sprint(args...), "err", err)

}