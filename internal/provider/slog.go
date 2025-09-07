package provider

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/viper"
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
	// Configure handler options based on configuration
	opts := &slog.HandlerOptions{
		AddSource: viper.GetBool("logger.add_source"),
		Level:     getLevelFromConfig(),
	}

	// Determine output writer
	writer := getOutputWriter()

	// Determine logger format from configuration
	var handler slog.Handler
	if viper.GetBool("logger.pretty") {
		handler = prettySlogJSONHandler(writer, opts)
	} else {
		handler = slog.NewJSONHandler(writer, opts)
	}

	// Add default attributes if configured
	if viper.IsSet("logger.attributes") {
		attributes := viper.GetStringMapString("logger.attributes")
		attrs := make([]slog.Attr, 0, len(attributes))

		for key, value := range attributes {
			attrs = append(attrs, slog.String(key, value))
		}

		handler = handler.WithAttrs(attrs)
	}

	logger := slog.New(handler)
	return &slogState{logger: logger}
}

// getLevelFromConfig returns the appropriate log level from configuration
func getLevelFromConfig() slog.Level {
	level := strings.ToLower(viper.GetString("logger.level"))

	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo // Default to info
	}
}

// getOutputWriter returns the appropriate writer based on configuration
func getOutputWriter() *os.File {
	output := strings.ToLower(viper.GetString("logger.output"))

	if output == "file" && viper.GetString("logger.file_path") != "" {
		filePath := viper.GetString("logger.file_path")
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			return file
		}
		// Fall back to stdout if file can't be opened
	}

	return os.Stdout
}

type prettySlogHandler struct {
	opts  *slog.HandlerOptions
	w     *os.File
	attrs []slog.Attr
	group string
}

var _ slog.Handler = (*prettySlogHandler)(nil)

func (h *prettySlogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.opts.Level.Level()
}

func (h *prettySlogHandler) Handle(ctx context.Context, record slog.Record) error {
	data := map[string]any{
		"level": record.Level.String(),
		"time":  record.Time.Format("2006-01-02T15:04:05.000Z07:00"),
		"msg":   record.Message,
	}
	if h.group != "" {
		data["group"] = h.group
	}

	record.Attrs(func(attr slog.Attr) bool {
		data[attr.Key] = attr.Value.Any()

		return true
	})

	if len(h.attrs) > 0 {
		for _, attr := range h.attrs {
			data[attr.Key] = attr.Value.Any()
		}
	}

	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	h.w.Write(b)
	h.w.Write([]byte("\n"))

	return nil
}

func (h *prettySlogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	h.attrs = append(h.attrs, attrs...)
	return h
}

func (h *prettySlogHandler) WithGroup(name string) slog.Handler {
	h.group = name
	return h
}

// prettySlogJSONHandler creates a pretty JSON handler for slog
func prettySlogJSONHandler(w *os.File, opts *slog.HandlerOptions) slog.Handler {
	return &prettySlogHandler{
		opts: opts,
		w:    w,
	}
}
