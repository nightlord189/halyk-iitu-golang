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
	user.Rename("ADMIN")
	fmt.Println(user)
	fmt.Println(user.Stringify())
}

func (u *User) Rename(name string) {
	u.Name = name
	fmt.Printf("\nuser after rename: %v\n", u)
}

func (u User) Stringify() string {
	return fmt.Sprintf("%d:%s", u.ID, u.Name)
}
