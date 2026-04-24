package logger

import (
	"log/slog"
	"os"
)

type Interface interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

func New() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}
