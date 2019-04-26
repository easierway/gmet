package main

import (
	"github.com/easierway/gmet"
	"time"
)

func main() {
	met := gmet.CreateGMetInstanceByDefault("seelog.xml")
	met.PeriodicallyFlush(time.Second)
	for range time.Tick(time.Millisecond) {
		met.Metric("A", "A", "A")
		met.Metric("B")
		met.Metric("A", "B", "C")
	}
}
