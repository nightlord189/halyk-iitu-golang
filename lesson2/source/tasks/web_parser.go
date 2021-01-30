package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var workersCount = 50

func main() {
	now := time.Now()
	var wg sync.WaitGroup
	restrictedChan := make(chan struct{}, 1)
	for i := 1; i <= workersCount; i++ {
		wg.Add(1)
		go makeHttp(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", i), &wg, restrictedChan)
	}
	wg.Wait()
	fmt.Println("time:", time.Since(now))
}

func makeHttp(url string, wg *sync.WaitGroup, ch chan struct{}) {
	fmt.Printf("making http-request to %s\n", url)
	ch <- struct{}{}
	resp, err := http.Get(url)
	<-ch
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	wg.Done()
}
