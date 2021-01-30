package main

import (
	"fmt"
)

func main() {
	exChan := make(chan int, 3)
	go produce(exChan)
	go produce(exChan)
	for val := range exChan {
		fmt.Println("read", val)
	}
	fmt.Println("end")
}

func produce(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("produced")
	}
	//close(ch)
	fmt.Println("closed")
}
