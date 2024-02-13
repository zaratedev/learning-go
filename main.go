package main

import "fmt"

func main() {
	// Operadores relacionales

	/*
		>
		<
		>=
		<=
		==
		!=
	*/

	edad := 30

	fmt.Println(edad)
	fmt.Println(edad > 5)
	fmt.Println(edad < 5)
	fmt.Println(edad >= 5)
	fmt.Println(edad <= 5)
	fmt.Println(edad == 30)
	fmt.Println(edad != 40)

	// Operadores logicos

	fmt.Println("Operadores Logicos")

	// && || !

	operacion := 20 == 20 && 30 == 30

	fmt.Println("Operacion con &&", operacion)

	operacion = 10 > 40 || 50 < 50

	fmt.Println("Operacion con ||", operacion)
}
