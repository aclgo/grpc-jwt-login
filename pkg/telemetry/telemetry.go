package telemetry

import (
	"github.com/aclgo/grpc-jwt/pkg/logger"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	"golang.org/x/net/trace"
)

type Provider struct {
	log        logger.Logger
	metric     metric.Meter
	tracer     trace.Trace
	propagator propagation.TextMapPropagator
}

func NewProvider(log logger.Logger, metric metric.Meter, trace trace.Trace, propagation propagation.TextMapPropagator) *Provider {
	return &Provider{
		log:        log,
		metric:     metric,
		tracer:     trace,
		propagator: propagation,
	}
}

func (p *Provider) Logger() logger.Logger {
	return p.log
}
func (p *Provider) Tracer() trace.Trace {
	return p.tracer
}
func (p *Provider) Meter() metric.Meter {
	return p.metric
}
func (p *Provider) Propagation() propagation.TextMapPropagator {
	return p.propagator
}
