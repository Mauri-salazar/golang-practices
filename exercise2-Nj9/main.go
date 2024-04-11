package main

import "fmt"

// Este ejercicio te ayudará a reforzar el concepto de method sets:
// Crea un tipo struct persona
// Adjunta el método hablar usando un receptor de tipo puntero
// *persona
// Crea un tipo interfaz humano
// Para implícitamente implementar esa interfaz, un tipo humano debe tener el método hablar
// Crea la función “diAlgo”
// Haz que tome un humano como parámetro
// Haz que llame al método hablar
// Muestra lo siguiente en tu código
// PUEDES pasar un valor de tipo *persona a diAlgo
// NO puedes pasar un valor de tipo persona a diAlgo

type person struct {
	name string
}

type human interface {
	talk()
}

// this a method func (receptor) name (parameters)
func (p *person) talk() {
	fmt.Println("Hello, my name is", p.name)
}

// this a method
func saySomeThing(h human) {
	h.talk()
}

func main() {
	p := person{
		name: "Mauricio",
	}
	saySomeThing(&p)
}
