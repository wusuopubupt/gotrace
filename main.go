package main

import (
	"fmt"
	"github.com/divan/gotrace/trace"
	"os"
	"strings"
	"time"
)

func main() {
	var src EventSource
	if len(os.Args) > 1 {
		src = NewNativeRun(os.Args[1])
	} else {
		src = NewTraceSource("trace.out", "trace.bin")
	}

	events, err := src.Events()
	if err != nil {
		panic(err)
	}
	c := &cmdWriter{}

	fmt.Printf("Got %d events\n", len(events))
	var lastG uint64
	for _, ev := range events {
		switch ev.Type {
		case trace.EvGoStart:
			if len(ev.Stk) > 0 {
				if strings.HasPrefix(ev.Stk[0].Fn, "runtime") {
					if ev.Stk[0].Fn != "runtime.main" {
						break
					}
				}
				c.StartGoroutine(ev.Ts, ev.Stk[0].Fn, ev.G, lastG)
			}
			lastG = ev.Args[0]
		case trace.EvGoEnd:
			c.StopGoroutine(ev.Ts, "", ev.G)
			lastG = 0
		//case trace.EvGoBlockSend:
		//	fmt.Printf("%v: Block Send %v\n", time.Duration(ev.Ts), ev.G)
		//case trace.EvGoBlockRecv:
		//	fmt.Printf("%v: Block Recv %v\n", time.Duration(ev.Ts), ev.G)
		case trace.EvGoSend:
			c.ChanSend(ev.Ts, ev.Args[1], ev.G, ev.Args[0])
		case trace.EvGoRecv:
			c.ChanRecv(ev.Ts, ev.Args[1], ev.G, ev.Args[0])
		case trace.EvGoBlockSelect:
			fmt.Printf("%v: Select %v\n", time.Duration(ev.Ts), ev.G)
			lastG = 0
		case trace.EvGoStop, trace.EvGoSched, trace.EvGoPreempt,
			trace.EvGoSleep, trace.EvGoBlock, trace.EvGoBlockSend, trace.EvGoBlockRecv,
			trace.EvGoBlockSync, trace.EvGoBlockCond, trace.EvGoBlockNet,
			trace.EvGoSysBlock:
			lastG = 1
		case trace.EvGoUnblock:
			lastG = ev.Args[0]
		}
	}
}
