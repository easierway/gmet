package gmet

import (
	"bytes"
	"fmt"
	"time"
)

type JSON_Formatter struct{}

const (
	SEC_SPLITTER               string = "\t"
	FIELD_SPLITTER             string = ":"
	TIMESTAMP_KEY              string = "timestamp"
	SEC_SPLITTER_REPLACEMENT   string = "_"
	FIELD_SPLITTER_REPLACEMENT string = "-"
)

func valueToJSON(v interface{}) string {
	switch v.(type) {
	case string:
		return "\"" + v.(string) + "\""
	case time.Time:
		return "\"" + (v.(time.Time)).Format(time.RFC3339Nano) + "\""
	default:
		return fmt.Sprintf("%v", v)
	}
}

func keyToJSON(k string) string {
	return "\"" + k + "\""
}

func toJSON_SEC(k string, v interface{}) string {
	return keyToJSON(k) + ":" + valueToJSON(v)
}

func (formatter *JSON_Formatter) Format(registry Registry) (string, error) {
	buf := bytes.NewBufferString("")
	buf.WriteString("{")
	buf.WriteString(toJSON_SEC(TIMESTAMP_KEY, time.Now()))
	// buf.WriteString(",")
	// buf.WriteString(toJSON_SEC(HOST_ADDR, HostAddr.Value))
	// buf.WriteString(",")
	// buf.WriteString(toJSON_SEC(HOST_NAME, HostName.Value))
	// buf.WriteString(",")
	// buf.WriteString(toJSON_SEC(SYSTYPE, SysType.Value))

	registry.Each(func(name string, i interface{}) {
		switch metric := i.(type) {
		case Counter:
			buf.WriteString(",")
			buf.WriteString(toJSON_SEC(name, metric.Count()))
		}
	})
	buf.WriteString("}")
	return buf.String(), nil
}
