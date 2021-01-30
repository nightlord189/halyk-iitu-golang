package main

import (
	"fmt"
	"sync"
)

type ConcurrentStruct struct {
	sync.RWMutex
	Map map[string]string
}

func main() {
	fmt.Println("start")
	cs := ConcurrentStruct{
		Map: make(map[string]string),
	}
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(3)
		go WriteParallel(&cs, 30, &wg)
		go ReadParallel(&cs, 30, &wg)
		go ReadParallel(&cs, 30, &wg)
	}
	wg.Wait()
	fmt.Println("end")
}

func WriteParallel(cs *ConcurrentStruct, count int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		cs.RWMutex.Lock()
		cs.Map["key1"] = fmt.Sprintf("value%d", i)
		cs.RWMutex.Unlock()
	}
}

func ReadParallel(cs *ConcurrentStruct, count int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		cs.RWMutex.RLock()
		fmt.Println(cs.Map["key1"])
		cs.RWMutex.RUnlock()
	}
}
