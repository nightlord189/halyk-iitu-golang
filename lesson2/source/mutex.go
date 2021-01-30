package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("start")
	mp := make(map[string]string)
	var mutex sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go WriteParallel(mp, 30, &wg, &mutex)
	}
	wg.Wait()
	fmt.Println("end")
}

func WriteParallel(mp map[string]string, count int, wg *sync.WaitGroup, mt *sync.Mutex) {
	defer wg.Done()
	defer mt.Unlock()
	for i := 0; i < count; i++ {
		mt.Lock()
		mp["key1"] = fmt.Sprintf("value%d", i)
		fmt.Println(mp["key1"])
		delete(mp, "key")
		mt.Unlock()
	}
}
