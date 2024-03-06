package main

import (
	"fmt"
)

type User struct {
	Name   string
	Email  string
	Active bool
}

type Student struct {
	User User
	Id   string
}

type Course struct {
	Title  string
	Videos []Video
}

type Video struct {
	Title  string
	Course Course
}

func main() {
	// Relacion uno a uno
	user1 := User{Name: "Jonathan", Email: "zaratedev@gmail.com", Active: true}
	user2 := User{Name: "Cynthia", Email: "cynthia@gmail.com", Active: false}

	student1 := Student{User: user1, Id: "1f1"}

	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Println(student1)

	// Acceder a datos del usuario desde student

	fmt.Println(student1.User.Email)

	// Relacion uno a muchos

	video1 := Video{Title: "Instalacion de Go"}
	video2 := Video{Title: "Introduccion a Go"}

	videos := []Video{video1, video2}

	curse1 := Course{Title: "Curso de Go", Videos: videos}

	video1.Course = curse1
	video2.Course = curse1

	fmt.Println(video1.Course.Title)
	fmt.Println(video2.Course.Title)
	fmt.Println("Lista de videos")
	for _, video := range curse1.Videos {
		fmt.Println(video.Title)
	}
}
