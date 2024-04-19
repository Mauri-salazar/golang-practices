package main

import "fmt"

func main() {
	//buffered  channel (con bufer)
	ch := make(chan int, 1)

	ch <- 45

	fmt.Println("Successful buffered", <-ch)
}
