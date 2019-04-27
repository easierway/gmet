package gmet

import (
	"net"
	"os"
	"time"
)

var HostAddr string
var HostName string
var SysType string

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

func CreateGMetInstanceByDefault(seelogCfg string, sysType string) GMetInstance {
	// create a metric writer
	writer, err := CreateMetWriterBySeeLog(seelogCfg)
	if err != nil {
		panic(err)
	}
	// create GMet instance by given the writer and the formatter
	gmet := CreateGMetInstance(&JSON_Formatter{}, writer)
	SysType = sysType
	return gmet
}

func init() {
	var err error
	HostAddr, err = IpAddress()
	if err != nil {
		HostAddr = err.Error()
	}
	hostname, err := os.Hostname()
	if err != nil {
		HostName = err.Error()
	} else {
		HostName = hostname
	}
	// default
	SysType = MISSING_VALUE
}

// Get the local IP address
func IpAddress() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return MISSING_VALUE, err
	}
	for _, address := range addrs {
		// Check if it is ip circle
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}

		}
	}
	return MISSING_VALUE, err
}
