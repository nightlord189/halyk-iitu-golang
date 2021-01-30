package main

import (
	"fmt"
	"time"
)

func main() {
	exChan := make(chan int)
	go produce(exChan)
	val := <-exChan
	fmt.Println(val)
}

func produce(ch chan int) {
	time.Sleep(3 * time.Second)
	ch <- 500
}
