package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type test struct {
		dates    []int
		response float64
	}

	tests := []test{
		test{[]int{1, 4, 6, 8, 100}, 6},
		test{[]int{0, 8, 10, 1000}, 9},
		test{[]int{9000, 4, 10, 8, 6, 12}, 9},
		test{[]int{123, 744, 140, 200}, 170},
	}

	for _, v := range tests {
		x := CenteredAvg(v.dates)

		if x != v.response {
			t.Error("Expected", v.response, "Got", x)
		}
	}

}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 4, 6, 8, 100}))

	//output
	//6
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 4, 6, 8, 100000})
	}
}
