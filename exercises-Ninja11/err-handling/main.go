package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	Name     string
	LastName string
	Phrases  []string
}

func main() {
	p1 := person{
		Name:     "James",
		LastName: "Bond",
		Phrases:  []string{"Shaken, not stirred", "any last whishes?, never say never."},
	}

	bs, err := json.Marshal(p1)

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(bs))
}
