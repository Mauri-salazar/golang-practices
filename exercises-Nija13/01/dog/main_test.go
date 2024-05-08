package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	y := Years(10)

	if y != 70 {
		t.Error("Expected", 70, "got", y)
	}
}

func ExampleYears() {
	fmt.Println(Years(10))

	//output
	//70
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func TestYearsTwo(t *testing.T) {
	y := YearsTwo(20)

	if y != 140 {
		t.Error("Expected", 140, "got", y)
	}
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(20))

	//output
	//140
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(20)
	}
}
