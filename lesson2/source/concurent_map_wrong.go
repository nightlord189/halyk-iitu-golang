package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("start")
	mp := make(map[string]string)
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go WriteParallel(mp, 30, &wg)
	}
	wg.Wait()
	fmt.Println("end")
}

func WriteParallel(mp map[string]string, count int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		mp["key1"] = fmt.Sprintf("value%d", i)
		fmt.Println(mp["key1"])
		delete(mp, "key")
	}
}
