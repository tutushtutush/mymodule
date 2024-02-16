package mypackage

import "testing"

func TestAdd(t *testing.T) {
	result := Add(3, 5)
	expected := 8
	if result != expected {
		t.Errorf("Add(3, 5) returned %d, expected %d", result, expected)
	}
}
