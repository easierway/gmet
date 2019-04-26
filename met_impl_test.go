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

func BenchmarkGMetInstance_Metric(b *testing.B) {
	gmet := CreateGMetInstance(&JSON_Formatter{}, &DummyWriter{})
	b.Run("metric A", func(b *testing.B) {
		gmet.Metric("A")
	})
	b.Run("metric A B C D", func(b *testing.B) {
		gmet.Metric("A", "B", "C", "D")
	})
	b.Run("metric A B C D * N", func(b *testing.B) {
		gmet.Metric("A", "B", "C", "D", "A", "B", "C", "D", "A", "B", "C", "D", "A", "B", "C", "D", "A", "B", "C", "D")
	})
}
