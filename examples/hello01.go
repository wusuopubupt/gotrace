package main

import "time"

func main() {
	// create new channel of type int
	ch := make(chan int)

	// start new anonymous goroutine
	go func() {
		time.Sleep(10 * time.Millisecond)
		// send 42 to channel
		ch <- 42
		time.Sleep(10 * time.Millisecond)
	}()
	// read from channel
	<-ch
}
