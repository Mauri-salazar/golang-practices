package main

import "fmt"

//select

func main() {
	out := make(chan int)
	c := gen(out)

	receiver(c, out)
	fmt.Println("About to end")
}

func gen(out chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i <= 10; i++ {
			c <- i
		}
		out <- 1 //when we're finish the to add values at channel c, we're to  add the value 1 at channel out.
		close(c)
	}()
	return c
}

func receiver(c, out <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("Printing from C", v)
		case <-out:
			return
		}
	}

}
