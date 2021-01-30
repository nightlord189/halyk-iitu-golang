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
	var mp sync.Map
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(3)
		go WriteParallel(&mp, 30, &wg)
		go ReadParallel(&mp, 30, &wg)
		go ReadParallel(&mp, 30, &wg)
	}
	wg.Wait()
	mp.Range(func(key, value interface{}) bool {
		fmt.Printf("\nkey: %v, value: %v", key, value)
		return true
	})
	fmt.Println("end")
}

func WriteParallel(cs *sync.Map, count int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		cs.Store("key1", fmt.Sprintf("value%d", i))
	}
}

func ReadParallel(cs *sync.Map, count int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		fmt.Println(cs.Load("key1"))
	}
}
