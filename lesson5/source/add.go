package main

import "fmt"

func Add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("start")
	a := 3
	b := 8
	c := Add(a, b)
	fmt.Printf("%d + %d = %d", a, b, c)
}
