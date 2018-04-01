package main

import (
	"os"
	"runtime/trace"
	"time"
	"fmt"
)

func main() {
	trace.Start(os.Stderr)

	i := 0
	for i = 0; i < 3; i++ {
		// start new anonymous goroutine
		go func() {
			time.Sleep(10 * time.Millisecond)
			fmt.Println("hello, world from goroutine")
			time.Sleep(10 * time.Millisecond)
		}()
	}

	time.Sleep(1 * time.Second)

	trace.Stop()
}
