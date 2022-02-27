package app

type Dependencies struct {
	GenerateID
}

type GenerateID func() (string, error)

type Logger interface {
	Info(args ...interface{})
	Error(err error, args ...interface{})
}