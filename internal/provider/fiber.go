package provider

import (
	"app/pkg/example"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Arsfiqball/codec/flame"
	"github.com/gofiber/contrib/otelfiber/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/trace"
)

type Fiber interface {
	App() *fiber.App
	Serve(ctx context.Context) error
	Clean(ctx context.Context) error
	Readiness(ctx context.Context) error
}

type fiberState struct {
	app  *fiber.App
	addr string
}

type FiberFeatureSet struct {
	Example *example.Example
}

func MakeFiber(fs FiberFeatureSet, ot Otel, lg Slog) (Fiber, error) {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			unpacked := flame.Unpack(err)
			httpUnpacked := flame.HttpUnpack(err)
			span := trace.SpanFromContext(ctx.UserContext())
			traceId := span.SpanContext().TraceID().String()
			spanId := span.SpanContext().SpanID().String()

			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			if httpUnpacked.Code >= 0 && httpUnpacked.Code < 600 {
				code = httpUnpacked.Code
			}

			type respT struct {
				Code    string `json:"code"`
				Info    string `json:"info"`
				TraceId string `json:"traceId"`
				SpanId  string `json:"spanId"`
				Data    any    `json:"data,omitempty"`
			}

			if ctx.Request().Header.Peek("Accept") != nil &&
				string(ctx.Request().Header.Peek("Accept")) == "application/json" {
				return ctx.Status(code).JSON(respT{
					Code:    unpacked.Code,
					Info:    httpUnpacked.Message,
					TraceId: traceId,
					SpanId:  spanId,
					Data:    httpUnpacked.Data,
				})
			}

			ctx.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)

			return ctx.Status(code).SendString(fmt.Sprintf(`
				<!DOCTYPE html>
				<html lang="en">
				<head>
					<meta charset="UTF-8">
					<meta name="viewport" content="width=device-width, initial-scale=1.0">
					<title>Error %d</title>
				</head>
				<body>
					<h1>Error %d</h1>
					<p>%s</p>
					<p>Trace ID: %s</p>
					<p>Span ID: %s</p>
					<p>Code: %s</p>
					<p>Data: %v</p>
				</body>
				</html>
			`, code, code, httpUnpacked.Message, traceId, spanId, unpacked.Code, httpUnpacked.Data))
		},
	})

	// Add OpenTelemetry middleware
	app.Use(otelfiber.Middleware(
		otelfiber.WithTracerProvider(ot.Provider()),
		otelfiber.WithNext(func(c *fiber.Ctx) bool {
			return c.Path() == "/readiness" // Skip OpenTelemetry for readiness endpoint
		}),
	))

	app.Use(func(c *fiber.Ctx) (err error) {
		start := time.Now()

		defer func() {
			duration := time.Since(start)
			span := trace.SpanFromContext(c.UserContext())
			traceId := span.SpanContext().TraceID().String()
			spanId := span.SpanContext().SpanID().String()

			if err != nil {
				unpacked := flame.Unpack(err)
				httpUnpacked := flame.HttpUnpack(err)

				if httpUnpacked.Code >= 500 {
					stack := flame.StackFrom(err, 10)

					lg.Logger().Error(httpUnpacked.Message,
						"method", c.Method(),
						"path", c.Path(),
						"code", unpacked.Code,
						"status", httpUnpacked.Code,
						"duration", duration,
						"remote_ip", c.IP(),
						"user_agent", c.Get("User-Agent"),
						"trace_id", traceId,
						"span_id", spanId,
						"stack", stack,
					)
				} else {
					lg.Logger().Warn(httpUnpacked.Message,
						"method", c.Method(),
						"path", c.Path(),
						"code", unpacked.Code,
						"status", httpUnpacked.Code,
						"duration", duration,
						"remote_ip", c.IP(),
						"user_agent", c.Get("User-Agent"),
						"trace_id", traceId,
						"span_id", spanId,
					)
				}

				return
			}

			// TODO: Add more context to the log if needed
			lg.Logger().Info(c.Response().String(),
				"info", "Request completed successfully",
				"method", c.Method(),
				"path", c.Path(),
				"status", c.Response().StatusCode(),
				"duration", duration,
				"remote_ip", c.IP(),
				"user_agent", c.Get("User-Agent"),
				"trace_id", traceId,
				"span_id", spanId,
			)
		}()

		err = c.Next()

		return
	})

	app.Use(func(c *fiber.Ctx) (err error) {
		defer flame.RecoverAs(&err, 10)

		return c.Next()
	})

	app.Get("/readiness", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(200)
	})

	app.Mount("/example", fs.Example.Fiber())

	app.Use(func(c *fiber.Ctx) error {
		// Handle 404 Not Found
		return flame.NotFound().Here()
	})

	addr := viper.GetString("serve.addr")
	if addr == "" {
		addr = ":3000"
	}

	return &fiberState{app: app, addr: addr}, nil
}

func (f *fiberState) App() *fiber.App {
	return f.app
}

func (f *fiberState) Serve(ctx context.Context) error {
	return f.app.Listen(f.addr)
}

func (f *fiberState) Clean(ctx context.Context) error {
	return f.app.Shutdown()
}

func (f *fiberState) Readiness(ctx context.Context) error {
	return nil
}

var ProvideFiber = wire.NewSet(
	MakeFiber,
	wire.Struct(new(FiberFeatureSet), "*"),
)
