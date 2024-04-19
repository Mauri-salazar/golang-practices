package main

import "fmt"

func main() {
	//buffered  channel (con bufer)
	ch := make(chan int, 2)

	ch <- 45
	ch <- 50

	fmt.Println("Successful buffered", <-ch)
	fmt.Println("Successful buffered", <-ch)
}
