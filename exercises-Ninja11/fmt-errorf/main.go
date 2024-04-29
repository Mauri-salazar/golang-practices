package main

import (
	"encoding/json"
	"errors"
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
		Phrases:  []string{"Shaken, not stirred", "any last wishes?", "Never say never."},
	}

	bs, err := aJSON(p1)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))

}

// aJSON also need to return an error
func aJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, errors.New(fmt.Sprintf("Error in aJSON: %v", err)) // We can using  fmt.Errorf() string, error

	}
	return bs, nil
}
