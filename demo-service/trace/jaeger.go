package trace

import (
	"github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

func NewTracer(serverName, addr string) (opentracing.Tracer, io.Closer, error) {
	configuration := jaegercfg.Configuration{
		ServiceName: serverName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			CollectorEndpoint:   "http://jaeger:14268/api/traces",
		},
	}
	sender, err := jaeger.NewUDPTransport(addr, 0)
	if err != nil {
		return nil, nil, err
	}
	reporter := jaeger.NewRemoteReporter(sender)
	return configuration.NewTracer(jaegercfg.Reporter(reporter))
}
