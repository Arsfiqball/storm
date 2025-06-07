package provider

import (
	"log/slog"
	"os"
)

type Slog interface {
	Logger() *slog.Logger
}

type slogState struct {
	logger *slog.Logger
}

func (s *slogState) Logger() *slog.Logger {
	return s.logger
}

func ProvideSlog() Slog {
	// Create a new slog logger with JSON format
	// and output to standard output.
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return &slogState{logger: logger}
}
