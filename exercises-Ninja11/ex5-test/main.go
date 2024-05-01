package main

/*
Una simple función que suma (y que vamos a probar más adelante)
en Go
@author parzibyte
*/

import "fmt"

func main() {
	resultado := Sumar(1, 2)
	fmt.Printf("El resultado de la suma es: %d", resultado)
}

func Sumar(a int, b int) int {
	return a + b
}
