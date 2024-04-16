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

	gr := 100

	var wg sync.WaitGroup
	wg.Add(gr)

	var mu sync.Mutex // I use Mutex for block the execution of Goroutines  with the methods Lock () and Unlock()

	for i := 0; i < gr; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			counter = v
			fmt.Println(counter)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}

// Them to know If the program have data rance condition I execute the command go run "-rance" main.go
// I use Mutex of package sync for delete the race condition of my  program.
