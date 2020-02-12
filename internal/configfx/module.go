package configfx

import (
	"github.com/yurykabanov/fxm/viperfx"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(viperfx.PFlagsProvider),
	fx.Provide(viperfx.MakeViperProvider(
		viperfx.WithEnvPrefix("SOME_SERVICE"),
		viperfx.WithConfigName("some_service"),
		viperfx.WithAdditionalConfigPaths("/etc/some_service"),
	)),
)
