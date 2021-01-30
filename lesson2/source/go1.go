package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("start")
	fmt.Printf("GOMAXPROCS %d", runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(4)
	go WriteParallel(10)
	go WriteParallel(10)
	go WriteParallel(10)
	time.Sleep(3 * time.Second)
	fmt.Println("end")
}

func WriteParallel(count int) {
	for i := 0; i < count; i++ {
		fmt.Printf("parallel %d\n", i)
		fmt.Printf("num of goroutines: %d\n", runtime.NumGoroutine())
	}
}
