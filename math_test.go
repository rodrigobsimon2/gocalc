package main

import "testing"

func TestSum(t *testing.T) {
	expected := 30

	if got := Sum(15, 15); got != expected {
		t.Errorf("expected %d, got %d", expected, got)
	}
}