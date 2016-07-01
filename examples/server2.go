package main

import (
	"fmt"
	"net"
	"os"
	"runtime/trace"
	"time"
)

func handler(c net.Conn, ch chan int) {
	ch <- len(c.RemoteAddr().String())
	c.Write([]byte("ok"))
	c.Close()
}

func logger(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func server(l net.Listener, ch chan int) {
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(c, ch)
	}
}

func main() {
	trace.Start(os.Stderr)

	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}
	ch := make(chan int)
	go logger(ch)
	go server(l, ch)
	time.Sleep(1 * time.Second)
	trace.Stop()
}
