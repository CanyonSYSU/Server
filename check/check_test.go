package check

import (
	"strconv"
	"testing"
)

type functionFilter func(string) (bool, error)

type inputAndExpected struct {
	input    string
	expected bool
}

func mTest(shouldSuccess []inputAndExpected, t *testing.T, testType string, filter functionFilter) {
	for _, ts := range shouldSuccess {
		if real, _ := filter(ts.input); real != ts.expected {
			expected_String := strconv.FormatBool(ts.expected)
			unexpected_String := strconv.FormatBool(!ts.expected)
			t.Errorf("fail: %s %s should be test %s but %s", testType, ts.input, expected_String, unexpected_String)
		} else {
			t.Log("success " + ts.input)
		}
	}
}

func Test_Check_Phone(t *testing.T) {
	shouldSuccess := []inputAndExpected{
		{"123456", false},
		{"12ab45", false},
		{"abcd", false},
		{"abcdefghijk", false},
		{"13719179988", true},
	}
	mTest(shouldSuccess, t, "number", CheckPhone)
}

func Test_Check_Name(t *testing.T) {
	shouldSuccess := []inputAndExpected{
		{"123", true},
		{"_@fad", false},
		{"你好", true},
	}
	mTest(shouldSuccess, t, "name", CheckName)
}

func Test_Check_Serve_Time(t *testing.T) {
	shouldSuccess := []inputAndExpected{
		{"25:00-12:00", false},
		{"12:00-12:61", false},
		{"8:00-10:00", false},
		{"05:61-25:11", false},
		{"12:00-15:30", true},
	}
	mTest(shouldSuccess, t, "serveTime", Checkservertime)
}
