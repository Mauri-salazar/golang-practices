package main

/*
Probar la función Sumar que existe en main.go
@author parzibyte
*/
// Importa el paquete:
import "testing"

// Escribe TestXXXX en donde XXXX es el nombre de la función original
func TestSumar(t *testing.T) { // Recibir struct de tipo testing.T
	resultado := Sumar(1, 2)
	resultadoEsperado := 3
	if resultado != resultadoEsperado {
		t.Errorf("Error en %s. Se esperaba %d pero el resultado fue %d", t.Name(), resultadoEsperado, resultado)
	}
	// Y si no, o sea, si todo va bien, no hacemos nada
}
