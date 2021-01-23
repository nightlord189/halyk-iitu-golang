package main

import (
	"fmt"
)

func main() {
	loadConfig("")
	fmt.Println("all is ok!")
}

func loadConfig(path string) {
	defer func() {
		val := recover()
		if val != nil {
			fmt.Printf("panic recovered: %v\n", val)
		}
	}()
	if path == "" {
		panic("PATH IS EMPTY!")
	}
}
