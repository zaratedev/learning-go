package main

import (
	"fmt"
)

func main() {
	// Condicionales
	var edad int
	// fmt.Print("Ingresa tu edad: ")
	// fmt.Scanf("%d", &edad)

	edad = 18

	if edad >= 18 {
		fmt.Println("Usuario mayor de edad")
	} else {
		fmt.Println("Usuario menor de edad")
	}

	// Multiples condiciones

	var calificacion int

	fmt.Print("Ingresa una calificacion: ")
	fmt.Scanf("%d", &calificacion)

	if calificacion == 10 {
		fmt.Println("Felicidades aprobaste la materia")
	} else if calificacion == 8 || calificacion == 9 {
		fmt.Println("Aprobaste la materia.")
	} else if calificacion == 6 || calificacion == 7 {
		fmt.Println("Aprobaste la materia. Pero necesitas estudiar mas")
	} else if calificacion >= 0 && calificacion <= 5 {
		fmt.Println("Reprobsate la materia")
	} else {
		fmt.Println("Calificación no valida")
	}

	// Usando switch
	switch calificacion {
	case 10:
		fmt.Println("Felicidades aprobaste la materia")
	case 8, 9:
		fmt.Println("Aprobaste la materia.")
	case 7, 6:
		fmt.Println("Aprobaste la materia. Pero necesitas estudiar mas")
	case 1, 2, 3, 4, 5:
		fmt.Println("Reprobsate la materia")
	default:
		fmt.Println("Calificación no valida")
	}
}
