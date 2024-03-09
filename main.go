package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	id       int
	username string
	email    string
	age      int
}

var id int
var users map[int]User

func createUser() {
	clear()
	fmt.Println("Ingresa un valor para username: ")
	username := readLine()

	fmt.Println("Ingresa un valor para email: ")
	email := readLine()

	fmt.Println("Ingresa un valor para edad: ")
	age, err := strconv.Atoi(readLine())

	if err != nil {
		panic("Valor para edad incorrecto")
	}

	id++

	user := User{id, username, email, age}

	users[id] = user

	fmt.Println("Usuario creado exitosamente")
}

func listUsers() {
	clear()

	for id, user := range users {
		fmt.Println(id, "-", user.username)
	}

	fmt.Println("\n")
}

func updateUser(values string...) {

}

func deleteUser() {
	clear()

	fmt.Println("Ingresa el id del usuario a eliminar")
	id, err := strconv.Atoi(readLine())

	if err != nil {
		panic("Valor incorrecto para ID")
	}

	if _, ok := users[id]; ok {
		delete(users, id)
	}

	fmt.Println("Usuario eliminado")
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func readLine() string {
	if option, err := reader.ReadString('\n'); err != nil {
		panic("No es posible obtener el valor")
	} else {
		return strings.TrimSuffix(option, "\n")
	}
}

func main() {
	var option string
	users = make(map[int]User)

	reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Println("A) Crear")
		fmt.Println("B) Listar")
		fmt.Println("C) Actualizar")
		fmt.Println("D) Eliminar")

		fmt.Print("Ingresa una opcion: ")

		option = readLine()

		if option == "quit" || option == "q" {
			break
		}

		switch option {
		case "a", "crear":
			createUser()
		case "b", "listar":
			listUsers()
		case "c", "actualizar":
			updateUser()
		case "d", "eliminar":
			deleteUser()
		default:
			fmt.Println("Opcion no encontrada")
		}
	}
	
	fmt.Println("Bye")
}
