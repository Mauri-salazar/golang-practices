package main

import "fmt"

func main() {
	c := make(chan int)    // General channel  This can send and receive dates between goroutines
	cs := make(chan<- int) // Send only channel This only send values
	cr := make(<-chan int) // Receive only channel This only receive values

	fmt.Println("------------")
	fmt.Printf("c\t%T/n", c)
	fmt.Printf("cs\t%T/n", cs)
	fmt.Printf("cr\t%T/n", cr)

	// Conversion of channels, We can to conversion the channels specifics to general "chan<- or <-chant type to chan type", but we
	// can't to conversion the channels generals to channels specifics EX: "chan type to <-chan type or cha<- type"
	fmt.Println("------------")
	fmt.Printf("c\t%T/n", (chan<- int)(c))
	fmt.Printf("c\t%T/n", (<-chan int)(c))

	//Also We can to assign a channels specifics to channels generals "<-chan or chan<- type = chan type",
	// but we can't do otherwise  "chan type = <-chan or chan<- type"

	cr = c
	cs = c
}
