package gmet

import (
	"time"
)

type GMetInstance struct {
	registry  Registry
	formatter MetFormatter // metrics formatter
	writer    MetWriter    // metrics data writer
}

func (gmet *GMetInstance) Metric(keys ...string) {
	for _, key := range keys {
		gmet.registry.GetOrRegister(key, NewCounter()).(Counter).Inc(1)
	}
}

// flush registry to writer
func (gmet *GMetInstance) Flush() {
	// replace with new one
	registry := gmet.registry
	gmet.registry = NewRegistry()

	if formatted, err := gmet.formatter.Format(registry); err != nil {
		return
	} else {
		gmet.writer.Write(formatted)
	}
}

// periodically flush
func (gmet *GMetInstance) PeriodicallyFlush(freq time.Duration) {
	go func() {
		for range time.Tick(freq) {
			gmet.Flush()
		}
	}()
}

func CreateGMetInstance(formatter MetFormatter, writer MetWriter) GMetInstance {
	ins := GMetInstance{registry: NewRegistry(), formatter: formatter, writer: writer}
	return ins
}

func CreateGMetInstanceByDefault(seelogCfg string) GMetInstance {
	// create a metric writer
	writer, err := CreateMetWriterBySeeLog(seelogCfg)
	if err != nil {
		panic(err)
	}
	// create GMet instance by given the writer and the formatter
	gmet := CreateGMetInstance(&JSON_Formatter{}, writer)
	return gmet
}
