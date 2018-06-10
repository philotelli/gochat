package trace

import (
	"fmt"
	"io"
)

//Tracer is the interface that describes an object capable of tracing events throughout code
type Trace interface {
	Trace(...interface{})
}

func New(w io.Writer) Trace {
	return &trace{out: w}
}

type trace struct {
	out io.Writer
}

func (t *trace) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Println(t.out)
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

// Off creates a Tracer that will ignore calls to Trace
func Off() Trace {
	return &nilTracer{}
}
