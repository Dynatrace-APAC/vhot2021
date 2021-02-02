package main

import (
	"go.opentelemetry.io/otel/api/global"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// initGlobalTracer sets up a basic OpenTelemetry configuration
func initGlobalTracer(providerOpts []sdktrace.TracerProviderOption) error {
	providerOpts = append(providerOpts, sdktrace.WithConfig(sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}))

	tp := sdktrace.NewTracerProvider(providerOpts...)
	global.SetTracerProvider(tp)

	return nil
}
