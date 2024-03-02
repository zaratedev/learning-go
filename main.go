package main

import (
	"errors"
	"fmt"
)

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("No se puede dividir por 0")
	}

	return a / b, nil
}

func main() {
	// Excepciones

	result, err := div(100, 10)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Resultado de la division", result)
	}
}
