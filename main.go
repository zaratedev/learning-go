package main

import (
	"fmt"
)

type Animal interface {
	// Definir metodos
	eat()
	sleep()
}

type Dog struct {
	Name string
}

func (dog Dog) eat() {
	fmt.Println(dog.Name, "Esta comiendo")
}

func (dog Dog) sleep() {
	fmt.Println(dog.Name, "Esta durmiendo")
}

func executeActions(animal Animal) {
	animal.eat()
	animal.sleep()
}

// Interfaces
func main() {
	dog := Dog{Name: "Docker"}

	fmt.Println(dog.Name)

	executeActions(dog)
}
