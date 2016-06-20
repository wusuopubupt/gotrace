package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		time.Sleep(10 * time.Millisecond)
		task, ok := <-ch
		if !ok {
			return
		}
		fmt.Println("processing task", task)
	}
}

func pool(wg *sync.WaitGroup, workers, tasks int) {
	ch := make(chan int)

	for i := 0; i < workers; i++ {
		time.Sleep(1 * time.Millisecond)
		go worker(ch, wg)
	}

	for i := 0; i < tasks; i++ {
		time.Sleep(10 * time.Millisecond)
		ch <- i
	}

	close(ch)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(36)
	go pool(&wg, 36, 50)
	wg.Wait()
}
