package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"time"
)

func main() {
	trace.Start(os.Stderr)
	out := make(chan int)
	quit := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(10 * time.Microsecond)
			v := <-out
			fmt.Println(v)
		}
	}()
	go func() {
		time.Sleep(1 * time.Second)
		<-quit
	}()

	for i := 0; i < 10; i++ {
		select {
		case out <- i:
		case quit <- true:
		}
	}
	trace.Stop()
}
