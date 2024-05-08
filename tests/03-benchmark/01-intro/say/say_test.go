package say

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("John")
	if s != "Welcome Dear John" {
		t.Error("Expected", "Welcome Dear John", "Got", s)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("John"))

	//output:
	// Welcome Dear John
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("John")
	}
}
