package main

import "time"

func main() {
	// create new channel of type int
	ch := make(chan int)

	// start new anonymous goroutine
	go func() {
		// send 42 to channel
		ch <- 42
	}()
	// read from channel
	time.Sleep(10 * time.Millisecond)
	<-ch
}
