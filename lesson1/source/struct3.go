package main

import (
	"fmt"
)

type User struct {
	ID      int
	Name    string
	Profile Profile
	Friends []int
}

type Profile struct {
	Email   string
	Address string
}

func main() {
	user := User{
		ID:   15,
		Name: "User1",
		Profile: Profile{
			Email:   "test@gmail.com",
			Address: "Almaty",
		},
		Friends: []int{1, 5, 6},
	}
	user.ID = 16
	fmt.Println(user)
	rename(&user)
	fmt.Println(user)
}

func rename(u *User) {
	u.Name = "John"
	u.Profile = Profile{
		Email:   "test1@gmail.com",
		Address: "Aktobe",
	}
	fmt.Printf("\nuser after rename: %v\n", u)
}
