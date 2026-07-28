package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := TestEvenOrOdd(10)
	if result != "even" {
		t.Errorf("expected even, acutal: %s", result)
	}
}