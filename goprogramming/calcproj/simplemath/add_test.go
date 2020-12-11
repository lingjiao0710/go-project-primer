package simplemath

import "testing"

func TestAdd(t *testing.T) {
	r := Add(10, 12)
	if r != 22 {
		t.Errorf("Add(10, 12) failed, got %d, expected 22", r)
	}
}
