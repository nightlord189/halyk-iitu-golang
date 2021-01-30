package main

import (
	"fmt"
	"time"
)

func main() {
	exChan1 := make(chan int, 2)
	exChan2 := make(chan int, 2)
	go produce(exChan1)
	go produce(exChan2)
	select {
	case val1 := <-exChan1:
		fmt.Println("read from channel1:", val1)
	case val2 := <-exChan2:
		fmt.Println("read from channel2:", val2)
	default:
		fmt.Println("nothing to read")
	}
	fmt.Println("end")
}

func produce(ch chan<- int) {
	time.Sleep(2 * time.Second)
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("produced")
	}
	fmt.Println("closed")
}
