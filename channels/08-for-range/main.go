package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c) // to make for loop we need close the channels with the close() function
	}()

	for v := range c {
		fmt.Println(v)
	}
}
