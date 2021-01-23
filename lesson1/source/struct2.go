package main

import (
	"fmt"
)

type User struct {
	ID   int
	Name string
}

func main() {
	user := User{
		ID:   15,
		Name: "User1",
	}
	user.ID = 16
	fmt.Println(user)
	rename(&user)
	fmt.Println(user)
}

func rename(u *User) {
	u.Name = "John"
	fmt.Printf("\nuser after rename: %v\n", u)
}
