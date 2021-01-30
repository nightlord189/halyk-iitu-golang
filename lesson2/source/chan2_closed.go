package main

import (
	"fmt"
)

func main() {
	exChan := make(chan int)
	go produce(exChan)
	val := <-exChan
	fmt.Println(val)
	val = <-exChan
	fmt.Println(val)
	val = <-exChan
	fmt.Println(val)
	val = <-exChan
	fmt.Println(val)
}

func produce(ch chan int) {
	ch <- 500
	fmt.Println("produced")
	close(ch)
	fmt.Println("closed")
}
