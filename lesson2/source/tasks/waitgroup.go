package main

import (
	"fmt"
	"sync"
)

type WaitGroup struct {
	sync.Mutex
	count     int
	closeChan chan struct{}
}

func (wg *WaitGroup) Add(count int) {
	wg.Mutex.Lock()
	wg.count += count
	wg.Mutex.Unlock()
}

func (wg *WaitGroup) Done() {
	wg.Mutex.Lock()
	wg.count--
	wg.Mutex.Unlock()
	if wg.count == 0 {
		wg.closeChan <- struct{}{}
	}
}

func (wg *WaitGroup) Wait() {
	<-wg.closeChan
}

func main() {
	wg := WaitGroup{
		closeChan: make(chan struct{}),
	}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go produce(10, &wg)
	}
	wg.Wait()
	fmt.Println("main end")
}

func produce(count int, wg *WaitGroup) {
	for i := 0; i < count; i++ {
		fmt.Printf("produce %d\n", i)
	}
	wg.Done()
}
