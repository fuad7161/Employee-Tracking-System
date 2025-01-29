package util

import (
	"testing"
)

func TestStringToInt(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"123", 123},
		{"0", 0},
		{"456789", 456789},
		{"98", 98},
		{"1", 1},
	}

	for _, test := range tests {
		result := StringToInt(test.input)
		if result != test.expected {
			t.Errorf("StringToInt(%q) = %d; expected %d", test.input, result, test.expected)
		}
	}
}
