package main

import (
	"fmt"
)

func main() {
	result, desc := subtract2(2, 3)
	fmt.Println(result)
	fmt.Println(desc)
}

func add(a, b int) (int, string) {
	result := a + b
	description := "all is ok"
	return result, description
}

func subtract(a, b int) (result int, description string) {
	result = a - b
	description = "all is ok"
	return result, description
}

func subtract2(a, b int) (result int, description string) {
	result = a - b
	return
}
