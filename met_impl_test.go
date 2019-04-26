package gmet

import (
	"testing"
)

func TestGMetInstance_Metric(t *testing.T) {
	gmet := CreateGMetInstance(&JSON_Formatter{}, &DummyWriter{})
	gmet.Metric("A")
	gmet.Metric("A")
	gmet.Metric("A")
	gmet.Metric("B")
	gmet.Flush()
	gmet.Metric("B")
	gmet.Metric("B")
	gmet.Metric("B")
	gmet.Flush()
}
