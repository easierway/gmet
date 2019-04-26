package gmet

import (
	"strings"
	"testing"
)

func TestFormatMetricItemWithNoSplitterInValue(t *testing.T) {
	formatter := LtrFormatter{}
	registry := NewRegistry()
	registry.GetOrRegister("K1", NewCounter()).(Counter).Inc(1)
	registry.GetOrRegister("K2", NewCounter()).(Counter).Inc(1)
	formatted, err := formatter.Format(registry)
	if err != nil {
		t.Errorf("failed to format %v", err)
	}
	expected := "K1" + FIELD_SPLITTER + "1" +
		SEC_SPLITTER + "K2" + FIELD_SPLITTER + "1"
	if !strings.Contains(formatted, expected) {
		t.Errorf("The formatted is %s, but the items is expected as %s\n",
			formatted, expected)
	}

	if !strings.Contains(formatted, TIMESTAMP_KEY) {
		t.Error("Timestamp is missing\n")
	}
}

func TestFormatMetricItemWithSplitterInValue(t *testing.T) {
	formatter := LtrFormatter{}
	registry := NewRegistry()
	registry.GetOrRegister("K1\t", NewCounter()).(Counter).Inc(1)
	registry.GetOrRegister("K2:G", NewCounter()).(Counter).Inc(1)
	formatted, err := formatter.Format(registry)
	if err != nil {
		t.Errorf("failed to format %v", err)
	}
	expected := "K1_" + FIELD_SPLITTER + "1" +
		SEC_SPLITTER + "K2-G" + FIELD_SPLITTER + "1"
	if !strings.Contains(formatted, expected) {
		t.Errorf("The formatted is %s, but the items is expected as %s\n",
			formatted, expected)
	}

	if !strings.Contains(formatted, TIMESTAMP_KEY) {
		t.Error("Timestamp is missing\n")
	}
}
