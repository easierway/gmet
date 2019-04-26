// The format is [Metric Name][Field Splitter][Metric Value][Section Splitter][Metric Name][Field Splitter][Metric Value]
// Created on 2018.5
package gmet

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type LtrFormatter struct{}

func replaceSplitterCharsInValue(value string) string {
	p1 := strings.Replace(value, SEC_SPLITTER, SEC_SPLITTER_REPLACEMENT, -1)
	p2 := strings.Replace(p1, FIELD_SPLITTER, FIELD_SPLITTER_REPLACEMENT, -1)
	return p2
}

func (formatter *LtrFormatter) Format(registry Registry) (string, error) {
	buf := bytes.NewBufferString("")
	buf.WriteString(TIMESTAMP_KEY)
	buf.WriteString(FIELD_SPLITTER)
	buf.WriteString(strconv.FormatInt(time.Now().Unix(), 10))
	// buf.WriteString(SEC_SPLITTER)
	// buf.WriteString(HostAddr.Key)
	// buf.WriteString(FIELD_SPLITTER)
	// buf.WriteString(replaceSplitterCharsInValue(HostAddr.Value.(string)))
	// buf.WriteString(SEC_SPLITTER)
	// buf.WriteString(HostName.Key)
	// buf.WriteString(FIELD_SPLITTER)
	// buf.WriteString(replaceSplitterCharsInValue(HostName.Value.(string)))
	// buf.WriteString(SEC_SPLITTER)
	// buf.WriteString(SysType.Key)
	// buf.WriteString(FIELD_SPLITTER)
	// buf.WriteString(replaceSplitterCharsInValue(SysType.Value.(string)))
	// buf.WriteString(SEC_SPLITTER)

	registry.Each(func(name string, iface interface{}) {
		switch metric := iface.(type) {
		case Counter:
			buf.WriteString(replaceSplitterCharsInValue(name))
			buf.WriteString(FIELD_SPLITTER)
			value := fmt.Sprintf("%v", metric.Count())
			buf.WriteString(replaceSplitterCharsInValue(value))
			buf.WriteString(SEC_SPLITTER)
		}
	})
	return buf.String(), nil
}
