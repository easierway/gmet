package gmet

import (
	"fmt"
	"testing"
	"time"
)

func checkErr(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Errorf("Expected: %s,actual value: %s", expected, actual)
	}
}

func TestToJSON(t *testing.T) {
	timeStr := "2018-05-17T21:36:40.8487098+08:00"
	tObj, _ := time.Parse(time.RFC3339Nano, timeStr)
	ret := valueToJSON(tObj)
	checkErr(t, fmt.Sprintf("\"%s\"", timeStr), ret)

	intValue := 10
	ret = valueToJSON(intValue)
	checkErr(t, fmt.Sprintf("%d", intValue), ret)

	strValue := "Hello"
	ret = valueToJSON(strValue)
	checkErr(t, fmt.Sprintf("\"%s\"", strValue), ret)

}
