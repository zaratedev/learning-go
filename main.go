package main

import "fmt"

type Operation func(a, b int) int

func createOperation(operation string) Operation {
	if operation == "sum" {
		return func(a, b int) int { return a + b }
	}

	return func(a, b int) int { return a * b }
}

func main() {
	// Retornar funciones

	sum := createOperation("sum")

	result := sum(10, 20)

	fmt.Println(result)
}
