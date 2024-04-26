package main

import (
	"fmt"
	"runtime"
)

// Writer a program that:
// Add 100 numbers to the channel
// Get numbers of the channel and print in the console

func main() {
	fmt.Println("Numbers of Goroutines", runtime.NumGoroutine())
	c := make(chan int)

	launchGr(c)
	printCh(c)
	fmt.Println("End of program")

}

func launchGr(c chan int) {
	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
	}
	fmt.Println("Numbers of Goroutines", runtime.NumGoroutine())

}

func printCh(c <-chan int) {
	for v := 0; v < 100; v++ {
		fmt.Println(v, <-c)
	}
}
