package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	cs := make(chan<- int)

	go func() {
		cs <- 42
	}()
	cs = c
	fmt.Println(<-c)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
