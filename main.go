package main

import "fmt"

func factorial(number int) int {
	if number == 1 {
		return 1
	}

	return factorial(number-1) * number
}

func main() {
	// Funciones recursiva

	result := factorial(6)

	fmt.Println("Factorial es", result)
}
