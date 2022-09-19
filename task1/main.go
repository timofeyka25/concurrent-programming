package main

import (
	"fmt"
	"sync"
	"time"
)

func thread(wg *sync.WaitGroup, waitTime time.Duration) {
	defer wg.Done()
	fmt.Printf("Thread with sleep time %s started\n", waitTime)
	time.Sleep(waitTime)
	fmt.Printf("Thread with sleep time %s closed\n", waitTime)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	go thread(&wg, 1*time.Second)
	go thread(&wg, 2*time.Second)
	go thread(&wg, 3*time.Second)
	go thread(&wg, 4*time.Second)
	defer wg.Wait()
}
