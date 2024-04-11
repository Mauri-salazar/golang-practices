package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPUs\t\t\t", runtime.NumCPU())           // I get the numbers CPUs
	fmt.Println("Goroutines\t\t", runtime.NumGoroutine()) // I get the numbers of goroutines in execution.

	wg.Add(2)
	go bar()
	go foo()

	fmt.Println("System operator\t\t", runtime.GOOS) // I get the system operator of my computer.
	fmt.Println("ARCH\t\t", runtime.GOARCH)          // I get the Architecture of system.

	fmt.Println("CPUs\t\t\t", runtime.NumCPU())           // I get the numbers CPUs
	fmt.Println("Goroutines\t\t", runtime.NumGoroutine()) // I get the numbers of goroutines in execution.

	wg.Wait()

	fmt.Println("Finish the programer")
	fmt.Println("CPUs\t\t\t", runtime.NumCPU())           // I get the numbers CPUs
	fmt.Println("Goroutines\t\t", runtime.NumGoroutine()) // I get the numbers of goroutines in execution.

}

func bar() {
	fmt.Println("Printing from bar")
	wg.Done()
}

func foo() {
	fmt.Println("Printing from foo")
	wg.Done()
}
