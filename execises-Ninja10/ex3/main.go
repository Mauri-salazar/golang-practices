package main

import (
	"fmt"
)

// gor... receive
func main() {
	c := gen()
	receive(c)
	fmt.Println("End of program")
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func gen() <-chan int {
	c := make(chan int)
	// We're send values a other channels and this is with goroutine transmitter and receive.

	// gor.. transmitter
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		//we are use close()  because, we're use a loop for range and if we do not close it, it's execution will continue and the channels will be blocked.
		close(c)
	}()
	return c
}
