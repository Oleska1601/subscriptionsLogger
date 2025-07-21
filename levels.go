package logger

import "log/slog"

// LevelFatal defines the custom log level for logging a fatal level.
const (
	LevelFatal = slog.Level(12)
)

var levelMap = map[string]slog.Level{
	"DEBUG": slog.LevelDebug,
	"INFO":  slog.LevelInfo,
	"WARN":  slog.LevelWarn,
	"ERROR": slog.LevelError,
	"FATAL": LevelFatal,
}

var levelNames = map[slog.Level]string{
	slog.LevelDebug: "DEBUG",
	slog.LevelInfo:  "INFO",
	slog.LevelWarn:  "WARN",
	slog.LevelError: "ERROR",
	LevelFatal:      "FATAL",
}
