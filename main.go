package main

import (
	"os"
)

func main() {
	var src EventSource
	if len(os.Args) == 2 {
		src = NewNativeRun(os.Args[1])
	} else if len(os.Args) == 3 {
		src = NewTraceSource(os.Args[1], os.Args[2])
	} else {
		src = NewTraceSource("trace.out", "trace.bin")
	}

	events, err := src.Events()
	if err != nil {
		panic(err)
	}

	json, err := ConvertEvents(events)
	if err != nil {
		panic(err)
	}

	StartServer(":2000", json)
}
