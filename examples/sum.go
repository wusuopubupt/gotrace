package main

import (
	"fmt"
	"os"
	"time"
	"runtime/trace"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	trace.Start(os.Stderr)
	time.Sleep(100 * time.Millisecond)

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	go func() {c <- 1}()
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	trace.Stop()
}
