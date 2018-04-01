package main

import (
	"os"
	"runtime/trace"
	"time"
	"fmt"
)

func main() {
	trace.Start(os.Stderr)
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
	
    ret := <-ch

    fmt.Println(ret)
	trace.Stop()
}
