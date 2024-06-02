package main

import "testing"

func TestEventOrOdd(t *testing.T) {
	result := EventOrOdd(10)
	if result != "Even" {
		t.Errorf("expected: Even, actual: %s", result)
	}
}
