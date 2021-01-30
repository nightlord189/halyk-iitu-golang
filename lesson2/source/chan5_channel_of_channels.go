package main

import (
	"fmt"
	"time"
)

func main() {
	exChan := make(chan chan int, 3)
	go produce(exChan)
	time.Sleep(3 * time.Second)
	fmt.Printf("channel's length: %d, capacity: %d", len(exChan), cap(exChan))
	for val := range exChan {
		fmt.Printf("read chan with capacity: %d\n", cap(val))
	}
	fmt.Println("end")
}

func produce(ch chan chan int) {
	for i := 0; i < 5; i++ {
		ch <- make(chan int, i)
		fmt.Println("produced")
	}
	close(ch)
	fmt.Println("closed")
}
