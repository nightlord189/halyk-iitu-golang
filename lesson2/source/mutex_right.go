package main

import (
	"fmt"
	"sync"
)

type ConcurrentStruct struct {
	sync.Mutex
	store map[string]string
}

func (cs *ConcurrentStruct) Get(key string) string {
	cs.Mutex.Lock()
	value := cs.store[key]
	cs.Mutex.Unlock()
	return value
}

func main() {
	fmt.Println("start")
	cs := ConcurrentStruct{
		Map: make(map[string]string),
	}
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go WriteParallel(&cs, 30, &wg)
	}
	wg.Wait()
	fmt.Println("end")
}

func WriteParallel(cs *ConcurrentStruct, count int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		cs.Mutex.Lock()
		cs.Map["key1"] = fmt.Sprintf("value%d", i)
		fmt.Println(cs.Map["key1"])
		delete(cs.Map, "key")
		cs.Mutex.Unlock()
	}
}
