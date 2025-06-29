package stringutil

import "testing"

func TestInt2Str(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{0, "0"},
		{1, "1"},
		{-1, "-1"},
		{123, "123"},
		{-123, "-123"},
		{1000, "1000"},
		{-1000, "-1000"},
	}

	for _, test := range tests {
		result := Int2Str(test.input)
		if result != test.expected {
			t.Errorf("Int2Str(%d) = %s; want %s", test.input, result, test.expected)
		}
	}
}
