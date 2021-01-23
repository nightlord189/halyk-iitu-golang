package main

import (
	"fmt"
)

func main() {
	b := 11
	if b > 10 {
		fmt.Println("more")
	} else {
		fmt.Println("less")
	}
	switch b {
	case 5, 6, 7:
		fmt.Println("b is 5!")
	case 10:
		fmt.Println("b is 10!")
	default:
		fmt.Println("unknown!")
	}
}
