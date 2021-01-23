package main

import (
	"fmt"
)

func main() {
	arr1 := [4]int{1, 2, 4, 5}
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	fmt.Println(arr1)
	fmt.Println(slice1)
	slice1 = append(slice1, slice2...)
	fmt.Println(slice1)
	fmt.Println(len(arr1))
}
