package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Operating system", runtime.GOOS)
	fmt.Println("System Architecture", runtime.GOARCH)

}
