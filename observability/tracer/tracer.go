package tracer

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

func InitTracer(jaegerURL, serviceName string) (trace.Tracer, error) {
	exporter, err := NewJaegerExporter(jaegerURL)
	if err != nil {
		return nil, err
	}

	tp, err := NewTracerProvider(exporter, serviceName)
	if err != nil {
		return nil, err
	}

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{}, propagation.Baggage{},
	))

	return tp.Tracer("main tracer"), nil
}
