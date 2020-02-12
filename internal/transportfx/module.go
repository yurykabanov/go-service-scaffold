package transportfx

import (
	"github.com/yurykabanov/fxm/httpfx"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(httpfx.MakeHttpServerConfigProvider("http", httpfx.WithDefaultAddr("0.0.0.0:8000"))),
	fx.Provide(httpfx.HttpServerProvider),
	fx.Provide(EchoServerProvider),
	fx.Invoke(RunServer),
)
