package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(11)
	if result != "odd" {
		t.Errorf("expected odd, acutal: %s", result)
	}
}