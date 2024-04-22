package main

import "fmt"

func main() {
	even := make(chan int) // Name of variable is par
	odd := make(chan int)  // Name of variable is impar
	out := make(chan int)  // Name of variable is salir

	go send(even, odd, out)

	receiver(even, odd, out)

	fmt.Println("finalized")
}

func send(e, o, out chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	out <- 0
}

func receiver(e, o, out <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("Printed from the channel even", v)
		case v := <-o:
			fmt.Println("Printed from the channel odd", v)
		case v := <-out:
			fmt.Println("Printed from the channel out", v)
			return
		}
	}
}
