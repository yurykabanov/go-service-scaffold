package tracingfx

import (
	"github.com/yurykabanov/fxm/jaegertracingfx"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(JaegerConfigurationProvider),
	jaegertracingfx.JaegerTracerOption,
	fx.Provide(EchoTraceConfigProvider),
	fx.Provide(fx.Annotated{Name: "jaegertracing", Target: JaegerTracingMiddleware}),
	// jaegertracingfx.RegisterTracerAsGlobalOption,
)
