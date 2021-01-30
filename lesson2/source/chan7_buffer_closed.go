package main

import (
	"fmt"
	"time"
)

func main() {
	exChan := make(chan int, 3)
	go produce(exChan)
	time.Sleep(3 * time.Second)
	fmt.Printf("channel's length: %d, capacity: %d\n", len(exChan), cap(exChan))
	for val := range exChan {
		fmt.Println("read", val)
	}
	fmt.Println("end")
}

func produce(ch chan int) {
	for i := 0; i < 2; i++ {
		ch <- i
		fmt.Println("produced")
	}
	close(ch)
	fmt.Println("closed")
}
