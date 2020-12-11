package simplemath

import "testing"

func TestSqrt(t *testing.T) {
	r := Sqrt(100)
	if r != 10 {
		t.Errorf("Sqrt(100) failed, got %d, expected 10", r)
	}
}
