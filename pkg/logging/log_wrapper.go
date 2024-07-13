package logging

import (
	"context"
	"fmt"
	"github.com/ahghazey/logging/pkg/constant"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type loggingWrapper struct {
	*zap.Logger
}

func (l *loggingWrapper) Debugf(msg string, args ...any) {
	l.Logger.Debug(fmt.Sprintf(msg, args...))
}
func (l *loggingWrapper) Infof(msg string, args ...any) {
	l.Logger.Info(fmt.Sprintf(msg, args...))
}
func (l *loggingWrapper) Errorf(msg string, args ...any) {
	l.Logger.Error(fmt.Sprintf(msg, args...))
}
func (l *loggingWrapper) Fatalf(msg string, args ...any) {
	l.Logger.Fatal(fmt.Sprintf(msg, args...))
}
func (l *loggingWrapper) Debug(msg string) {
	l.Logger.Debug(msg)
}
func (l *loggingWrapper) Info(msg string) {
	l.Logger.Info(msg)
}
func (l *loggingWrapper) Error(msg string) {
	l.Logger.Error(msg)
}
func (l *loggingWrapper) Fatal(msg string) {
	l.Logger.Fatal(msg)
}
func (l *loggingWrapper) WithFields(fields map[string]string) (logger Logger) {
	if len(fields) == 0 {
		logger = &loggingWrapper{
			Logger: l.Logger,
		}
		return
	}
	zapFields := make([]zapcore.Field, 0)
	for k, v := range fields {
		zapFields = append(zapFields, zap.String(k, v))
	}

	clonedLog := l.Logger.With(zapFields...)
	logger = &loggingWrapper{
		Logger: clonedLog,
	}
	return
}
func (l *loggingWrapper) WithContext(ctx context.Context) Logger {
	if ctx == nil {
		return l
	}

	logger := l.Logger
	fields := []zap.Field{}

	if val, ok := ctx.Value(constant.RequestIdHeader).(string); ok {
		fields = append(fields, zap.String(string(constant.RequestIdHeader), val))
	}

	if val, ok := ctx.Value(constant.SpanIdHeader).(string); ok {
		fields = append(fields, zap.String(string(constant.SpanIdHeader), val))
	}

	if val, ok := ctx.Value(constant.TraceIdHeader).(string); ok {
		fields = append(fields, zap.String(string(constant.TraceIdHeader), val))
	}

	if val, ok := ctx.Value(constant.UserIdHeader).(string); ok {
		fields = append(fields, zap.String(string(constant.UserIdHeader), val))
	}

	return &loggingWrapper{
		Logger: logger.With(fields...),
	}
}

func newLogger(logger *zap.Logger) *loggingWrapper {

	return &loggingWrapper{
		Logger: logger,
	}

}
