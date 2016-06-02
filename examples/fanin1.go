package main

import (
	"fmt"
	"time"
)

func producer(ch chan int, d time.Duration) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(d)
	}
}

func reader(out chan int) {
	for x := range out {
		fmt.Println(x)
	}
}

func main() {
	ch := make(chan int)
	out := make(chan int)
	go producer(ch, 10*time.Millisecond)
	go producer(ch, 25*time.Millisecond)
	go reader(out)
	for i := range ch {
		out <- i
	}
}
