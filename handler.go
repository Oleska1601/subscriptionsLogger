package logger

import (
	"log/slog"
	"os"
)

func newHandler(level slog.Level) slog.Handler {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				return slog.String(slog.LevelKey, levelNames[level])
			}
			if a.Key == slog.MessageKey {
				return slog.Attr{Key: "path", Value: a.Value}
			}
			return a
		},
	})
	return handler
}
