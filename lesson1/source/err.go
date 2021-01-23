package main

import (
	"errors"
	"fmt"
)

func main() {
	result, _ := add(-1, 5)
	fmt.Println(result)
}

func add(a, b int) (int, error) {
	if a < 0 {
		return 0, errors.New("negative number")
	}
	return a + b, nil
}
