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
// The first return value is the fibonacci number to be calculated
// The second return value reports the number of recursive invocation that were required in order to calculate the result
func (f *fibonacci) Calc(n int) (int, int) {
	var result int
	var numIterations int

	if n < 3 {
		result = 1
		numIterations = 1
	} else {
		resultA, numIterationsA := f.Calc(n - 1)
		resultB, numIterationsB := f.Calc(n - 2)
		result = resultA + resultB
		numIterations = numIterationsA + numIterationsB + 1
	}
	return result, numIterations
}

/* LESSON 02: CHILD SPANS AND SPAN ATTRIBUTES

// Fibonacci calculates fibonacci numbers
type Fibonacci interface {
	Calc(n int) (int, int) // Calc has no documentation
}

// New creates a new Fibonacci Calculator
func New(ctx context.Context) Fibonacci {
	return &fibonacci{Context: ctx}
}

type fibonacci struct {
	Context context.Context
}

// Calc calculates the n-th fibonacci number
// The first return value is the fibonacci number to be calculated
// The second return value reports the number of recursive invocation that were required in order to calculate the result
func (f *fibonacci) Calc(n int) (int, int) {
	var span trace.Span
	tracer := global.Tracer("")
	f.Context, span = tracer.Start(f.Context, fmt.Sprintf("fib(%d)", n))
	defer span.End()

	var result int
	var numIterations int
	if n < 3 {
		result = 1
		numIterations = 1
	} else {
		resultA, numIterationsA := f.Calc(n - 1)
		resultB, numIterationsB := f.Calc(n - 2)
		result = resultA + resultB
		numIterations = numIterationsA + numIterationsB + 1
	}
	span.SetAttributes(label.Key("fib.result").Int(result))
	return result, numIterations
}
*/

/* LESSON 04: CREATING INSTRUMENTATION LIBRARIES

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
*/

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
