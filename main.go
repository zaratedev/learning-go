package main

import "fmt"

func main() {
	// Slices
	numbers := []int{1, 2, 3, 4} // Referencia a un arreglo base

	numbers = append(numbers, 5)
	numbers = append(numbers, 6)
	numbers = append(numbers, 7)
	numbers = append(numbers, 8)
	numbers = append(numbers, 9)
	numbers = append(numbers, 10)

	extract := numbers[0:5] // un nuevo slice a partir de uno

	numbers[0] = 100 // Aqui se modifica la posicion 0

	fmt.Println(numbers)
	fmt.Println(extract)

	// Parte dos

	months := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio"}

	// Puntero
	// Longitud
	// Capacidad

	longitud := len(months)
	capacidad := cap(months)

	fmt.Printf("Longitud %v y capacidad %v %p\n", longitud, capacidad, months)

	// Agregamos un nuevo mes
	months = append(months, "Julio")

	longitud = len(months)
	capacidad = cap(months)

	fmt.Printf("Longitud %v y capacidad %v %p\n", longitud, capacidad, months)
}
