package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("start")
	var once sync.Once
	for true {
		fmt.Printf("Enter command: ")
		var text string
		fmt.Scan(&text)
		if text == "signal" {
			once.Do(sendSignal)
		} else {
			fmt.Println("wrong command")
		}
	}
}

func sendSignal() {
	fmt.Println("signal sent")
}
