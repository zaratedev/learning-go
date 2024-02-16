package main

import (
	"fmt"
)

func main() {
	// Funcion Make

	slice := make([]int, 3, 3)

	slice[0] = 100
	slice[1] = 200
	slice[2] = 300

	slice = append(slice, 400)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// Mapas

	days := make(map[int]string)

	days[0] = "Domingo"
	days[1] = "Lunes"
	days[2] = "Martes"
	days[3] = "Miercoles"
	days[4] = "Jueves"
	days[5] = "Viernes"
	days[6] = "Sabado"

	// Eliminar un elemento del mapa
	delete(days, 4)

	fmt.Println(days)

	// Otro ejemplo
	students := make(map[string][]int)

	students["Jonathan"] = []int{100, 70, 90, 80, 50}

	fmt.Println(students)

	// Iterar un mapa

	users := map[int]string{} // Sin make
	users[1] = "Jonathan"
	users[2] = "Cynthia"
	users[3] = "Erick"
	users[4] = "Jorge"

	for id, name := range users {
		fmt.Println(id, name)
	}
}
