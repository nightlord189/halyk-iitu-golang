package main

import (
	"fmt"
)

func main() {
	defer closeFile()
	fmt.Println("file opened")
	defer fmt.Println("defer1")
	fmt.Println("processed")
	defer fmt.Println("defer2")
	fmt.Println("end")
	b := 1
	if b == 1 {
		return
	}
	defer fmt.Println("defer3")
}

func closeFile() {
	fmt.Println("file closed")
}
