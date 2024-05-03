package main

import (
	"fmt"
)

type dog struct {
	name string
	age  int
}

func main() {
	d1 := dog{
		name: "Tom",
		age:  dog.yearsDog(2),
	}
	fmt.Println(d1)
}
