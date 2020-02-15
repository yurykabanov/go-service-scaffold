package transportfx

import (
	"github.com/yurykabanov/fxm/httpfx"
	"github.com/yurykabanov/service-scaffold/pkg/transport/http/controller"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(httpfx.MakeHttpServerConfigProvider("http", httpfx.WithDefaultAddr("0.0.0.0:8000"))),
	fx.Provide(httpfx.HttpServerProvider),
	fx.Provide(EchoServerProvider),
	fx.Invoke(RunServer),

	fx.Provide(controller.NewVersionController),
)
