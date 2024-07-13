package logging

import (
	"context"
	"fmt"
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"sync"
)

var LogHandle AppLogger
var once sync.Once

type AppLogger struct {
	logger *loggingWrapper
}

func InitLogger(lvl, serviceName, environment string) error {
	level, err := parseLevel(lvl)
	LogHandle = getLogger(level, serviceName, environment)
	return err
}

func getLogger(level zapcore.Level, serviceName string, environment string) AppLogger {
	once.Do(func() {
		encoderConfig := ecszap.EncoderConfig{
			EncodeName:     zap.NewProductionEncoderConfig().EncodeName,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeDuration: zapcore.MillisDurationEncoder,
			EncodeCaller:   ecszap.FullCallerEncoder,
		}
		core := ecszap.NewCore(encoderConfig, os.Stdout, level)
		l := zap.New(core, zap.AddCaller())
		l = l.With(zap.String("app", serviceName)).With(zap.String("env", environment))

		zapLogger := newLogger(l)

		LogHandle = AppLogger{
			logger: zapLogger,
		}
	})

	return LogHandle
}

func (l *AppLogger) Debugf(msg string, args ...any) {
	l.logger.Debugf(msg, args...)
}
func (l *AppLogger) Infof(msg string, args ...any) {
	l.logger.Infof(msg, args...)
}
func (l *AppLogger) Errorf(msg string, args ...any) {
	l.logger.Errorf(msg, args...)
}
func (l *AppLogger) Fatalf(msg string, args ...any) {
	l.logger.Fatalf(msg, args...)
}
func (l *AppLogger) Debug(msg string) {
	l.logger.Debug(msg)
}
func (l *AppLogger) Info(msg string) {
	l.logger.Info(msg)
}
func (l *AppLogger) Error(msg string) {
	l.logger.Error(msg)
}
func (l *AppLogger) Fatal(msg string) {
	l.logger.Fatal(msg)
}
func (l *AppLogger) WithFields(fields map[string]string) Logger {
	return l.logger.WithFields(fields)
}
func (l *AppLogger) WithContext(ctx context.Context) Logger {
	return l.logger.WithContext(ctx)
}

func parseLevel(lvl string) (zapcore.Level, error) {
	switch strings.ToLower(lvl) {
	case "debug":
		return zap.DebugLevel, nil
	case "info":
		return zap.InfoLevel, nil
	case "warn":
		return zap.WarnLevel, nil
	case "error":
		return zap.ErrorLevel, nil
	}
	return zap.InfoLevel, fmt.Errorf("invalid log level <%v>", lvl)
}
