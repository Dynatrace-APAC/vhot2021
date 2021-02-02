package fibonacci

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/trace"
	"go.opentelemetry.io/otel/label"
)

// Fibonacci calculates fibonacci numbers
type Fibonacci interface {
	Calc(n int) (int, int) // Calc has no documentation
}

// New creates a new Fibonacci Calculator
func New() Fibonacci {
	return &fibonacci{}
}

type fibonacci struct{}

// Calc calculates the n-th fibonacci number
func (f *fibonacci) Calc(n int) (int, int) {
	if n < 3 {
		return 1, 1
	}
	a, iterA := f.Calc(n - 1)
	b, iterB := f.Calc(n - 2)
	return a + b, iterA + iterB + 1
}

// Wrap produces a Fibonacci Calculator with tracing capabilities
func Wrap(ctx context.Context, f Fibonacci) Fibonacci {
	return &tracingFib{Context: ctx, Fibonacci: f}
}

type tracingFib struct {
	Fibonacci Fibonacci
	Context   context.Context
}

func (tf *tracingFib) Calc(n int) (int, int) {
	var span trace.Span
	tracer := global.Tracer("")
	tf.Context, span = tracer.Start(tf.Context, fmt.Sprintf("fib(%d)", n))
	defer span.End()
	result, iterations := tf.Fibonacci.Calc(n)
	span.SetAttributes(label.Key("fib.result").Int(result))
	return result, iterations
}

/*
DO NOT REMOVE ANY TEXT BELOW THIS LINE






















*/

// PreserveImport is never getting called.
// In order to keep this session as simple as possible it's purpose is to preserve imports on top of this file
func PreserveImport() {
	lbl := label.Key("")
	tracer := global.Tracer("")
	var ctx context.Context
	var span trace.Span
	fmt.Printf("%v, %v, %v, %v", span, ctx, tracer, lbl)
}
