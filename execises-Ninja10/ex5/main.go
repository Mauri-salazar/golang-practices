package main

import "fmt"

// v, ok, check if a value exist, if exist return the value and bool true and if not exist return value 0 and bool false

func main() {
	c := make(chan int)

	go func() {
		c <- 2
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
