package main

import (
	"fmt"
	"os"
)

func function1() {
	fmt.Println("Function 1")
}

func function2() {
	fmt.Println("Function 2")
}

func main() {
	// Progrmar funciones

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Ha ocurrido un error")
		}
	}()

	if file, err := os.Open("samples.txt"); err != nil {
		panic("No es posible obtener el archivo")
	} else {
		defer func() {
			fmt.Println("close file")
			file.Close()
		}()

		content := make([]byte, 254)
		length, _ := file.Read(content)

		body := string(content[0:length])

		fmt.Println(body)
	}
}
