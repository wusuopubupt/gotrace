package main

import "time"

func main() {
	time.Sleep(100 * time.Millisecond)
	ch := make(chan int)
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch <- 42
		time.Sleep(100 * time.Millisecond)
	}()
	<-ch
	time.Sleep(100 * time.Millisecond)
}
