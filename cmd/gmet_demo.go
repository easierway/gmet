package main

import (
	. "github.com/easierway/gmet"
	"time"
)

func main() {
	gmet := CreateGMetInstance(&JSON_Formatter{})
	gmet.PeriodicallyFlush(time.Duration(10e8))
	for range time.Tick(time.Duration(10e7)) {
		gmet.Metric("A")
		gmet.Metric("B")
		gmet.Metric("A")
		gmet.Metric("A")
	}
}
