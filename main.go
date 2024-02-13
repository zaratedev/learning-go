package main

import "fmt"

func main() {
	// Trabajando con strings

	curso := "Curso profesional de go"

	longitud := len(curso)

	fmt.Println(curso, longitud)

	caracter := curso[0]

	fmt.Println(caracter)
	fmt.Printf("%c", caracter)
}
