package main

import "fmt"

func main() {
	//unbuffered channel (canal sin bufer)
	ch := make(chan int)

	go func() {
		ch <- 45

	}()
	fmt.Println("This channel is unblocked", <-ch)
}
