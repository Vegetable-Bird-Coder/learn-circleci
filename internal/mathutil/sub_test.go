package mathutil

import "testing"

func TestSub(t *testing.T) {
	result := Sub(2, 3)
	expected := -1
	if result != expected {
		t.Errorf("Sub(2, 3) = %d; want %d", result, expected)
	}
}
