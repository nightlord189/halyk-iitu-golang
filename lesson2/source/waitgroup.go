package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("start")
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go WriteParallel(30, &wg)
	}
	wg.Wait()
	fmt.Println("end")
}

func WriteParallel(count int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		fmt.Printf("parallel %d\n", i)
	}
}
