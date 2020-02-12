package tracingfx

import (
	"strings"

	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
)

func EchoTraceConfigProvider(tracer opentracing.Tracer) *jaegertracing.TraceConfig {
	return &jaegertracing.TraceConfig{
		Tracer:  tracer,
		Skipper: func(ctx echo.Context) bool {
			if strings.HasPrefix(ctx.Request().URL.Path, "/swagger/") {
				return true
			}

			// other cases

			return false
		},
	}
}

func JaegerTracingMiddleware(config *jaegertracing.TraceConfig) echo.MiddlewareFunc {
	return jaegertracing.TraceWithConfig(*config)
}
