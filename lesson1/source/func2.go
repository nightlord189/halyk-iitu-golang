package main

import (
	"fmt"
)

func main() {
	fmt.Println(summarize(1, 10, 5))
}

func summarize(numbers ...int, numbers2 ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}
