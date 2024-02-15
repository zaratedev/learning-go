package main

import "fmt"

func main() {
	// Arreglos

	var numeros [5]int

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	numeros[3] = 40
	numeros[4] = 50

	fmt.Println(numeros)

	// Otra forma de declarar arreglos

	products := [5]string{"sofa", "cama", "silla", "lampara", "mesa"}

	fmt.Println(products)

	// Otra forma sin usar la cantidad

	prices := [...]int{200, 300, 400, 500}

	fmt.Println(prices)

	currency := [...]string{1: "MXN", 0: "USD", 2: "EUR"}

	fmt.Println(currency)

	sub_currency := currency[0:2] // Esto ya es un slice

	fmt.Println(sub_currency)
}
