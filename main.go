package main

import "fmt"

type UserType int

const (
	Teacher UserType = 1
	Student UserType = 2
)

type User struct {
	Username string
	Type     UserType
}

func main() {
	user1 := User{Username: "zaratedev", Type: Student}
	user2 := User{Username: "Jonathan", Type: Teacher}

	fmt.Println(user1)
	fmt.Println(user2)
}
