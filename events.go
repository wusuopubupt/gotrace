package main

import (
	"github.com/divan/gotrace/trace"
	"os"
)

// EventSource defines anything that can
// return tracing events.
type EventSource interface {
	Events() ([]*trace.Event, error)
}

// TraceSource implements EventSource for
// pre-made trace file
type TraceSource struct {
	// Trace is the path to the trace file.
	Trace string

	// Binary is a path to binary, needed to symbolize stacks.
	// In the future it may be dropped, see
	// https://groups.google.com/d/topic/golang-dev/PGX1H8IbhFU
	Binary string
}

// NewTraceSource inits new TraceSource.
func NewTraceSource(trace, binary string) *TraceSource {
	return &TraceSource{
		Trace:  trace,
		Binary: binary,
	}
}

// ReadEvents reads trace file from filesystem and symbolizes
// stacks.
func (t *TraceSource) ReadEvents() ([]*trace.Event, error) {
	f, err := os.Open(t.Trace)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	events, err := trace.Parse(f)
	if err != nil {
		return nil, err
	}

	err = trace.Symbolize(events, t.Binary)

	return events, err
}

type NativeRun struct{}
type DockerRun struct{}
