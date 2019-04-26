package main

import (
	"github.com/easierway/gmet"
	"time"
)

func main() {
	met := gmet.CreateGMetInstanceByDefault("seelog.xml")
	met.PeriodicallyFlush(time.Duration(10e8))
	for range time.Tick(time.Duration(10e7)) {
		met.Metric("A", "A", "A")
		met.Metric("B")
		met.Metric("A", "B", "C")
	}
}
