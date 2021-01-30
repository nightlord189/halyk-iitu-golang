package main

import (
	"fmt"
	"time"
)

func main() {
	signalChan := make(chan struct{})
	exChan1 := make(chan int, 2)
	exChan2 := make(chan int, 2)
	go produce(exChan1)
	go produce(exChan2)
	go produceStopSignal(signalChan)
	run := true
	for run {
		select {
		case val1 := <-exChan1:
			fmt.Println("read from channel1:", val1)
		case val2 := <-exChan2:
			fmt.Println("read from channel2:", val2)
		case <-time.After(10 * time.Second):
			fmt.Println("timeout")
			run = false
		case <-signalChan:
			fmt.Println("close")
			run = false
		}
	}
	fmt.Println("end")
}

func produceStopSignal(ch chan<- struct{}) {
	time.Sleep(500 * time.Millisecond)
	ch <- struct{}{}
}

func produce(ch chan<- int) {
	time.Sleep(5 * time.Second)
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("produced")
	}
	fmt.Println("closed")
}
