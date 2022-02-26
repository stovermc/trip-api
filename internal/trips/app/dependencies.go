package app

type Logger interface {
	Info(args ...interface{})
	Error(err error, args ...interface{})
}