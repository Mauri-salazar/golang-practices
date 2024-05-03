package main

import "testing"

func TestSum(t *testing.T) {
	v := sum(2, 8)

	if v != 10 {
		t.Error("Expected", 10, "Got", v)
	}
}
