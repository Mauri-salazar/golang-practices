package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs", runtime.NumCPU())                       // I get the numbers CPUs
	fmt.Println("Numbers of Goroutine", runtime.NumGoroutine()) // I get the numbers Goroutine

	counter := 0

	const gr = 100

	var wg sync.WaitGroup
	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			v := counter
			v++
			runtime.Gosched()
			counter = v
			wg.Done()
			fmt.Println("Numbers of Goroutines", runtime.NumGoroutine())
		}()
	}
	fmt.Println("Counter:", counter)
	wg.Wait()
}

// Them to know If the program have data rance condition I execute the command go run "-rance" main.go
