package main

import (
	"fmt"
	"time"
)

func main() {
	N := 1
	ch := make(chan int, N)

	go func(ch chan int) {
		for i := 0; i < N; i++ {
			ch <- i
		}
		close(ch)
	}(ch)

	for v := range ch {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(v)
	}
}
