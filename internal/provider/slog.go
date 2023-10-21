package provider

import (
	"log/slog"
	"os"
)

func ProvideSlog() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}
