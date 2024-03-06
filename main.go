package main

import (
	"codigofacilito/plataform"
	"fmt"
)

func main() {
	course := plataform.Course{Title: "Curso de Go"}

	fmt.Println(course.GetTitle())
}
