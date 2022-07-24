package tracer

import (
	"fmt"

	servicename "github.com/NpoolPlatform/service-template/pkg/servicename"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func TraceID(span trace.Span, id string) trace.Span {
	span.SetAttributes(attribute.String("ID", id))
	return span
}

func TraceInvoker(span trace.Span, entity, module, invokeName string) trace.Span {
	span.AddEvent(fmt.Sprintf("%v.%v.%v.%v", servicename.ServiceName, entity, module, invokeName))
	return span
}

func TraceOffsetLimit(span trace.Span, offset, limit int) trace.Span {
	span.SetAttributes(
		attribute.Int("Offset", offset),
		attribute.Int("Limit", limit),
	)
	return span
}
