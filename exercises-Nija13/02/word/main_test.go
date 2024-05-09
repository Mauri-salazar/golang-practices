package word

import (
	"02/02/quote"
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	v := Count(quote.SunAlso)

	if v != 7251 {
		t.Error("Expected", 7251, "Got", v)
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func ExampleCount() {
	fmt.Println(Count(quote.SunAlso))

	//output
	//7251
}

func TestUseCount(t *testing.T) {
	m := UseCount("one two three three three")

	for k, v := range m {
		switch k {
		case "one":
			if v != 1 {
				t.Error("Expected", 1, "Got", v)
			}
		case "two":
			if v != 1 {
				t.Error("Expected", 1, "Got", v)
			}
		case "three":
			if v != 3 {
				t.Error("Expected", 3, "Got", v)
			}
		}
	}
}
