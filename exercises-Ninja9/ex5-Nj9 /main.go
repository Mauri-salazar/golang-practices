package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs", runtime.NumCPU())                       // I get the numbers CPUs
	fmt.Println("Numbers of Goroutine", runtime.NumGoroutine()) // I get the numbers Goroutine

	var counter int64 // I'm setting the value of the counter to the type int64(0)

	gr := 100

	var wg sync.WaitGroup
	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)            // I use AddInt64() that is a method in the atomic to increase a value.
			fmt.Println(atomic.LoadInt64(&counter)) // LoadInt64() print the value
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}

// Them to know If the program have data rance condition I execute the command go run "-rance" main.go
// I use Mutex of package sync for delete the race condition of my  program.
