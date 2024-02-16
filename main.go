package main

import (
	"fmt"
)

func main() {
	// Ciclo for

	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// For como while
	number := 12345
	count := 0

	for number > 0 {
		number = number / 10
		count++
	}

	fmt.Println("Cantidad de digitos", count)
}
