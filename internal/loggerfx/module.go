package loggerfx

import (
	"github.com/yurykabanov/fxm/logrusfx"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(logrusfx.MakeLoggerConfigProvider("log")),
	logrusfx.DefaultLoggersOption,
)
