package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)

	go func(ch chan int) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}(ch)

	for v := range ch {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(v)
	}
}
