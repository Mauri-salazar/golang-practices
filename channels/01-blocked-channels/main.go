package main

import "fmt"

func main() {
	//unbuffered channel (canal sin bufer)
	ch := make(chan int)

	ch <- 45
	fmt.Println("This channel is blocked", <-ch)
}
