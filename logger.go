package subscriptionsLogger

import (
	"context"
	"log/slog"
	"strings"
)

type LoggerInterface interface {
	Debug(string, string, ...any)
	Info(string, int, ...any)
	Warn(string, string, ...any)
	Error(string, string, int, error, ...any)
	Fatal(string, string, error, ...any)
}

type Logger struct {
	logger *slog.Logger
}

func New(inputLevel string) *Logger {
	level, ok := levelMap[strings.ToUpper(inputLevel)]
	if !ok {
		level = slog.LevelInfo
	}
	handler := newHandler(level)
	logger := slog.New(handler)
	return &Logger{logger: logger}
}

func (l *Logger) Debug(funcName string, methodName string, args ...any) {
	l.logger.Debug(funcName+" "+methodName, args...)
}

func (l *Logger) Info(funcName string, status int, args ...any) {
	l.logger.Info(funcName, append([]any{slog.Int("status", status)}, args...)...)
}

func (l *Logger) Warn(funcName string, methodName string, args ...any) {
	l.logger.Warn(funcName+" "+methodName, args...)
}

func (l *Logger) Error(funcName string, methodName string, status int, err error, args ...any) {
	l.logger.Error(funcName, append([]any{slog.Int("status", status), slog.Any("error", err)}, args...)...)
}

func (l *Logger) Fatal(funcName string, methodName string, err error, args ...any) {
	l.logger.Log(context.Background(), LevelFatal, funcName+" "+methodName, append([]any{slog.Any("error", err)}, args...)...)
}
