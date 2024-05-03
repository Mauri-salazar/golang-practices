package main

import (
	yearsDog "dog/dog"
	"fmt"
)

type dog struct {
	name string
	age  int
}

func main() {
	d1 := dog{
		name: "Tom",
		age:  yearsDog.YearsDog(2),
	}
	fmt.Println(d1)
}
