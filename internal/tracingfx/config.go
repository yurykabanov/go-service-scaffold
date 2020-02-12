package tracingfx

import (
	"time"

	jaegerconfig "github.com/uber/jaeger-client-go/config"
)

func JaegerConfigurationProvider() (*jaegerconfig.Configuration, error) {
	defaultConfig := jaegerconfig.Configuration{
		ServiceName: "some-service",
		Sampler: &jaegerconfig.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jaegerconfig.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}

	config, err := defaultConfig.FromEnv()
	if err != nil {
		return nil, err
	}

	return config, nil
}
