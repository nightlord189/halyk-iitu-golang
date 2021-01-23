package main

import (
	"fmt"
)

type User struct {
	ID    int
	Name  string
	Email *string
}

type Admin struct {
	User
	Rights string
	Secret string
}

func main() {
	user := Admin{
		Rights: "686",
		Secret: "_some_secret",
		User: User{
			ID:   15,
			Name: "User1",
		},
	}
	fmt.Println(user)
}
