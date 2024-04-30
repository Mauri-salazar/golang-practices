package main

import "fmt"

type errorPer struct {
	info string
}

func (err errorPer) Error() string {
	return fmt.Sprintf("Here's the error %v", err.info)
}

func main() {
	v := errorPer{
		info: "Error in function main",
	}

	foo(v)
}

func foo(err error) {
	fmt.Println("Run foo -", err, "\n", err.(errorPer).info) //In this case we using assertion err.(errorPer).info.
}
