package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"runtime/trace"
	"time"
)

var reqcount int64

func randomHandler(w http.ResponseWriter, req *http.Request) {
	buffered := bufio.NewWriter(w)
	w.Header().Set("Content-Length", "20000")
	for i := 0; i < 10000; i++ {
		fmt.Fprintln(buffered, rand.Int63n(10))
	}
	buffered.Flush()
}

func main() {
	trace.Start(os.Stderr)
	http.HandleFunc("/", randomHandler)
	go http.ListenAndServe(":5000", nil)
	time.Sleep(1 * time.Second)
	trace.Stop()
}
