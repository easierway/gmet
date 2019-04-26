package gmet

import (
	"log"
	"time"
)

const (
	HOST_ADDR     = "host"
	HOST_NAME     = "hostname"
	MISSING_VALUE = "N/A"
	SYSTYPE       = "systype"
)

type GMetInstance struct {
	registry  Registry
	formatter MetFormatter // metrics formatter
	// metWriter    MetWriter    // metrics data writer
	// mutex sync.RWMutex
}

func (gmet *GMetInstance) Metric(keys ...string) {
	for _, key := range keys {
		gmet.registry.GetOrRegister(key, NewCounter()).(Counter).Inc(1)
	}
}

func (gmet *GMetInstance) Flush() {
	// renew registry
	registry := gmet.registry
	gmet.registry = NewRegistry()

	if formatted, err := gmet.formatter.Format(registry); err != nil {
		return
	} else {
		// TODO: change to writer
		log.Println(formatted)
	}
}

func (gmet *GMetInstance) PeriodicallyFlush(freq time.Duration) {
	go func() {
		for range time.Tick(freq) {
			gmet.Flush()
		}
	}()
}

func CreateGMetInstance(formatter MetFormatter) GMetInstance {
	ins := GMetInstance{registry: NewRegistry(), formatter: formatter}
	return ins
}

func CreateGMetInstanceByDefault() GMetInstance {
	// create a metric writer
	// writer, err := CreateMetWriterBySeeLog(metricsFile)
	// if err != nil {
	// 	panic(err)
	// }
	// create GMet instance by given the writer and the formatter
	gmet := CreateGMetInstance(&JSON_Formatter{})
	return gmet
}
