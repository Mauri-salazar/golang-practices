package main

import "testing"

func TestFib(t *testing.T) {
	tables := []struct {
		a int
		n int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range tables {
		fib := Fibonacci(item.a)
		if fib != item.n {
			t.Error("Expect", fib, "Got", item.n)
		}
	}
}
