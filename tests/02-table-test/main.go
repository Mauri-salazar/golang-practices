package main

import "fmt"

func main() {
	fmt.Println("The sum of 2 + 4 is", sum(2, 4))
	fmt.Println("The sum of 1 + 5 is", sum(1, 5))
	fmt.Println("The sum of 6 + 7 is", sum(6, 7))
}

func sum(xi ...int) int {
	var sum int

	for _, v := range xi {
		sum += v
	}
	return sum
}
