package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func(ch chan int) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}(ch)

	for v := range ch {
		fmt.Println(v)
		time.Sleep(10 * time.Millisecond)
	}
}
