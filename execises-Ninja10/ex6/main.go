package main

import "fmt"

// Writer a program that:
// Add 100 numbers to the channel
// Get numbers of the channel and print in the console

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i <= 100; i++ {
			c <- i
		}
		close(c)
		return
	}()
	printCh(c)
	fmt.Println("End of program")

}

func printCh(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
