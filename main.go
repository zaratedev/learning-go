package main

import "fmt"

func main() {
	// Lectura de valores

	var name string
	var city string
	var height float32

	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &name)

	fmt.Print("Ingresa tu ciudad: ")
	fmt.Scanf("%s", &city)

	fmt.Print("Ingresa tu altura: ")
	fmt.Scanf("%f", &height)

	fmt.Printf("Hola %s tu ciudad es %s y tu altura %.2f\n", name, city, height)
}
