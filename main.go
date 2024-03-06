package main

import (
	"fmt"
)

type User struct {
	name  string
	email string
}

// Metodo

func (user *User) setName(name string) {
	user.name = name
}

func (user *User) getName() string {
	return user.name
}

func main() {
	// Estructuras

	var user User // Objeto

	user.name = "Jonathan Zarate"
	user.email = "zaratedev@gmail.com"

	fmt.Println(user.name)
	fmt.Println(user.email)

	person := User{name: "Cynthia", email: "cynthia@test.com"}

	fmt.Println(person)

	me := User{name: "Joe", email: "joe@gmail.com"}

	me.setName("Joe Doe")

	fmt.Println(me.getName())
}
