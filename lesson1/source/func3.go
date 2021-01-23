package main

import (
	"fmt"
)

func main() {
	call(add)
	call(subtract)
	func2 := func(a, b int) int {
		return 50
	}
	fmt.Println(func2(1, 5))
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func call(inputFunc func(int, int) int) {
	fmt.Println("Call from other func")
	fmt.Println(inputFunc(2, 1))
}
