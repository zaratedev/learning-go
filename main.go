package main

import "fmt"

func main() {
	name := "Jonathan"

	fmt.Println(name)

	// Multiples variables

	/**
	var lastname string
	var country string
	*/

	// Declaracion sin asignacion
	// var lastname, country string

	// Declaracion asignado valores
	// var lastname, country = "Zarate Hernandez", "Mexico"

	// Usando otra sintaxis
	lastname, country := "Zarate Hernandez", "Mexico"

	fmt.Println(lastname, country)
}
