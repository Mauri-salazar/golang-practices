package main

import (
	"errors"
	"fmt"
	"log"
)

type rootError struct {
	lat  string
	long string
	err  error
}

func (re rootError) Error() string {
	return fmt.Sprintf("Math error: %v %v %v", re.lat, re.long, re.err)
}

func main() {
	_, err := root(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func root(f float64) (float64, error) {
	if f < 0 {
		// Write your code here
		e := errors.New("I need to sleep more")
		return 0, rootError{
			lat:  "50.2289 N",
			long: "99.4656 W",
			err:  e,
		}
	}
	return 42, nil
}
