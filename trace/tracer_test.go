package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	const GREETING = "Hello trace package."
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Return from New should not be nil")
	} else {
		tracer.Trace(GREETING)
		if buf.String() != GREETING+"\n" {
			t.Errorf("Trace should not write '%s'.", buf.String())
		}
	}
}

func TestOff(t *testing.T) {
	var silentTracer Tracer = Off()
	silentTracer.Trace("something")
}
