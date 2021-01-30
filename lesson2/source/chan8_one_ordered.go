package main

import (
	"fmt"
)

func main() {
	exChan := make(chan int, 3)
	go produce(exChan)
	consume(exChan)
	fmt.Println("end")
}

func transformChannelToWrite(ch chan int) chan<- int {
	return ch
}

func transformChannelToRead(ch chan int) <-chan int {
	return ch
}

func produce(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("produced")
	}
	close(ch)
	fmt.Println("closed")
}

func consume(ch <-chan int) {
	for i := 0; i < 10; i++ {
		val := <-ch
		fmt.Println("read", val)
	}
}
