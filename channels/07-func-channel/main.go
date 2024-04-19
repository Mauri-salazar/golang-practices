package main

import "fmt"

func main() {
	c := make(chan int)

	go send(c)
	receive(c)
	fmt.Println("Finish")
}

func send(c chan<- int) {
	c <- 45
}

// In this case, I don't need that this func be one  goroutine because his code program ends with "fmt..."
func receive(c <-chan int) {
	fmt.Println(<-c)
}
