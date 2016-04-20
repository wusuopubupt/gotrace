package main

import (
	"bytes"
	"strings"

	"github.com/divan/gotrace/trace"
)

func ConvertEvents(events []*trace.Event) ([]byte, error) {
	var buf bytes.Buffer
	c := NewCmdWriter(&buf)

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
		case trace.EvGoSend:
			c.ChanSend(ev.Ts, ev.Args[1], ev.G, ev.Args[0])
		case trace.EvGoRecv:
			c.ChanRecv(ev.Ts, ev.Args[1], ev.G, ev.Args[0])
		case trace.EvGoBlockSelect:
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

	return buf.Bytes(), nil
}
