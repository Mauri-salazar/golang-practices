package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	type test struct {
		dates  []int
		result int
	}

	tets := []test{
		test{[]int{2, 4, 6}, 12},
		test{[]int{1, 5, 2}, 8},
		test{[]int{0, -1, 1}, 0},
		test{[]int{-10, 4, 20}, 14},
	}

	for _, v := range tets {
		x := sum(v.dates...)
		if x != v.result {
			t.Error("Expected", v.result, "Got", v)
		}
	}

}
