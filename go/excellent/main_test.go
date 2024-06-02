package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EventOrOdd(10)
	if result != "Even" {
		t.Errorf("expected: Even, actual: %s", result)
	}
}
