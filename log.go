package gmet

import (
	"time"
)

type Logger interface {
	Printf(format string, v ...interface{})
}

func Log(r Registry, freq time.Duration, l Logger) {
	LogScaled(r, freq, l)
}

// Output each metric in the given registry periodically using the given
// logger. Print timings in `scale` units (eg time.Millisecond) rather than nanos.
func LogScaled(r Registry, freq time.Duration, l Logger) {
	for range time.Tick(freq) {
		r.Each(func(name string, i interface{}) {
			switch metric := i.(type) {
			case Counter:
				l.Printf("counter %s\n", name)
				l.Printf("  count:       %9d\n", metric.Count())
			}
		})
	}
}

// Output each metric in the given registry periodically using the given
// logger. Print timings in `scale` units (eg time.Millisecond) rather than nanos.
func LogOnce(r Registry, l Logger) {
		r.Each(func(name string, i interface{}) {
			switch metric := i.(type) {
			case Counter:
				l.Printf("counter %s\n", name)
				l.Printf("  count:       %9d\n", metric.Count())
			}
		})
}
