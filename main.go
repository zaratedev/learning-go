package main

import "fmt"

type Operacion func(balance, cantidad int) int

func sum(a, b int) int { // Funcion de tipo operacion
	return a + b
}

func minus(a, b int) int { // Funcion de tipo operacion
	return a - b
}

func executeOperation(function Operacion) {
	result := function(10, 10)

	fmt.Println("Resultado:", result)
}

func main() {
	// Funciones como argumentos

	executeOperation(sum)
	executeOperation(minus)
}
