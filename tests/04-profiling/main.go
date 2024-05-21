// It allows us to know the part od the code we must improve for optimizer the program.
//commands
// create new file go test -cpuprofile=c.out
// run profiling file with go tool pprof c.out
// top show a little information
// list name of func, show a list information
//web create a  scheme and the open in the browser
// pdf save in a file pdf the scheme created  with the command web

package main

import "fmt"

func main() {
	fmt.Println("The sum of 2 + 4 is", Fibonacci(1))

}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
