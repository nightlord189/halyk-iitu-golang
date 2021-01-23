package main

import (
	"fmt"
)

var dict3 map[string]float64 = map[string]float64{"key1": 0.567}

func main() {
	//var dict1 map[string]float64 = map[string]float64{"key2": 0.567}
	var dict2 map[string]float64
	dict2 = make(map[string]float64)
	dict2["key67"] = 3.14
	fmt.Println(dict2)
	if _, ok := dict2["key68"]; ok {
		fmt.Println("key exists")
	} else {
		fmt.Println("key doesn't exist")
	}
}
