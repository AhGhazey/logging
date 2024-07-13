package logging

import "context"

type Logger interface {
	Debugf(msg string, args ...any)
	Infof(msg string, args ...any)
	Errorf(msg string, args ...any)
	Fatalf(msg string, args ...any)
	Debug(msg string)
	Info(msg string)
	Error(msg string)
	Fatal(msg string)
	WithFields(map[string]string) Logger
	WithContext(ctx context.Context) Logger
}
