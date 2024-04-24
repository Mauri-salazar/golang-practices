package main

import "fmt"

// In this case we are using a literal function because a channel without buffer,can't save values  to itself
func main() {
	c := make(chan int)
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)
}

//We are using the buffer

// func main() {
// 	c := make(chan int, 1)  ---- this is the buffer, add in the ends a second parameter, that is the numbers of values that can save
// 	c <- 45
// 	fmt.Println(<-c)
// }
