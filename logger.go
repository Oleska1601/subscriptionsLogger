package subscriptionsLogger

import (
	"context"
	"log/slog"
	"os"
	"strings"
)

type LoggerInterface interface {
	Debug(string, string, string, ...any)
	InfowithStatus(string, string, int, ...any)
	Info(string, string, ...any)
	Warn(string, string, string, ...any)
	ErrorWithStatus(string, string, string, int, error, ...any)
	Error(string, string, string, error, ...any)
	Fatal(string, string, string, error, ...any)
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

func (l *Logger) Debug(funcName string, actionName string, msg string, args ...any) {
	l.logger.Debug(msg, append([]any{slog.String("path", funcName+" "+actionName)}, args...)...)
}

func (l *Logger) InfowithStatus(funcName string, msg string, status int, args ...any) {
	l.logger.Info(msg, append([]any{slog.String("path", funcName), slog.Int("status", status)}, args...)...)
}

func (l *Logger) Info(funcName string, msg string, args ...any) {
	l.logger.Info(msg, append([]any{slog.String("path", funcName)}, args...)...)
}

func (l *Logger) Warn(funcName string, actionName string, msg string, args ...any) {
	l.logger.Warn(msg, append([]any{slog.String("path", funcName+" "+actionName)}, args...)...)
}

func (l *Logger) ErrorWithStatus(funcName string, actionName string, msg string, status int, err error, args ...any) {
	l.logger.Error(funcName, append([]any{slog.String("path", funcName+" "+actionName), slog.Int("status", status), slog.Any("error", err)}, args...)...)
}

func (l *Logger) Error(funcName string, actionName string, msg string, err error, args ...any) {
	l.logger.Error(funcName, append([]any{slog.String("path", funcName+" "+actionName), slog.Any("error", err)}, args...)...)
}

func (l *Logger) Fatal(funcName string, actionName string, msg string, err error, args ...any) {

	l.logger.Log(context.Background(), LevelFatal, msg, append([]any{slog.String("path", funcName+" "+actionName), slog.Any("error", err)}, args...)...)
	os.Exit(1)
}
