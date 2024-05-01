package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup // *waitGroup = add(), done() & wiat()

func main() {
	fmt.Println("Desde main")
	go bar()
	go foo()
	wg.Add(2) //contador
	wg.Wait() // revisa el contalor
}

func bar() {
	fmt.Println("Desde bar")
	wg.Done() // resta al contador 1
}
func foo() {
	fmt.Println("Desde foo")
	wg.Done()
}
