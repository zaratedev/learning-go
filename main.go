package main

import (
	"fmt"
)

func sum(a, b int) (int, string) {
	return a + b, "Sum sucess"
}

func main() {
	// Funciones

	operation, message := sum(50, 30)

	fmt.Println(operation, message)

	result, _ := sum(39, 49)

	fmt.Println(result)

	// Funciones anonimas

	func() {
		fmt.Println("Fucntion anonymous")
	}()

	myFunction := func(name string) string {
		message := fmt.Sprintf("Hola, %s", name)

		return message
	}

	message = myFunction("Jonathan")

	fmt.Println(message)
}
