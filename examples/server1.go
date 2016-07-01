package main

import (
	"net"
	"os"
	"runtime/trace"
	"time"
)

func handler(c net.Conn) {
	c.Write([]byte("ok"))
	c.Close()
}

func main() {
	trace.Start(os.Stderr)

	go time.AfterFunc(1*time.Second, func() { trace.Stop(); os.Exit(0) })

	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(c)
	}
}
