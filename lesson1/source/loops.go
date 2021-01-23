package main

import (
	"fmt"
)

func main() {
	arr1 := []int{1, 2, 5, 10, 15}
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}
	counter := 1
	fmt.Println()
	for ; counter < len(arr1); counter++ {
		fmt.Println(arr1[counter])
	}
	arr2 := []string{"Almaty", "Nur-Sultan", "Aktobe"}
	for _, value := range arr2 {
		fmt.Printf("\nindex: , value: %s\n", value)
	}
	for index := range arr2 {
		fmt.Println(index)
	}

	map1 := map[string]int{"key1": 15, "key2": 17, "key3": 19}
	for key, value := range map1 {
		fmt.Printf("\nindex: %s, value: %v\n", key, value)
	}
	for key := range map1 {
		fmt.Println(key)
	}
	stop := false
	for !stop {
		fmt.Println("running")
		stop = true
	}
}
