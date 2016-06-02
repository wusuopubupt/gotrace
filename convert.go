package main

import (
	"container/list"
	"fmt"
	"sort"
	"strings"

	"github.com/divan/gotrace/trace"
)

func ConvertEvents(events []*trace.Event) (Commands, error) {
	var c Commands

	sends := list.New()

	var lastG uint64
	for _, ev := range events {
		switch ev.Type {
		case trace.EvGoStart:
			if len(ev.Stk) > 0 {
				if strings.HasPrefix(ev.Stk[0].Fn, "runtime") {
					if ev.Stk[0].Fn != "runtime.main" {
						lastG = ev.Args[0]
						break
					}
				}
				c.StartGoroutine(ev.Ts, ev.Stk[0].Fn, ev.G, lastG)
			}
		case trace.EvGoCreate:
			lastG = ev.G
		case trace.EvGoEnd:
			c.StopGoroutine(ev.Ts, "", ev.G)
			lastG = 1
		case trace.EvGoSend:
			sends.PushBack(ev)
		case trace.EvGoRecv:
			send := findSource(sends, ev)
			if send == nil {
				fmt.Println("[WARN] Recv w/o Send:", ev)
				continue
			}
			c.ChanSend(send.Ts, ev.Args[1], send.G, ev.G, send.Args[2])
		case trace.EvGCStart, trace.EvGCDone, trace.EvGCScanStart, trace.EvGCScanDone:
			lastG = 1
		case trace.EvGoSched, trace.EvGoPreempt,
			trace.EvGoSleep, trace.EvGoBlock, trace.EvGoBlockSelect, trace.EvGoBlockSend, trace.EvGoBlockRecv,
			trace.EvGoBlockSync, trace.EvGoBlockCond, trace.EvGoBlockNet,
			trace.EvGoSysBlock:
			lastG = 1
		case trace.EvGoStop:
			lastG = 1
		}
	}

	// sort events
	sort.Sort(ByTimestamp(c))

	// insert stop main
	// TODO: figure out why it's not in the trace
	lastTs := c[len(c)-1].Time
	c.StopGoroutine(lastTs+1000, "", 1)

	return c, nil
}

// findSource tries to find corresponding Send event to ev.
func findSource(sends *list.List, ev *trace.Event) *trace.Event {
	for e := sends.Back(); e != nil; e = e.Prev() {
		send := e.Value.(*trace.Event)
		if send.Args[1] == ev.Args[1] && send.Args[0] == ev.Args[0] {
			sends.Remove(e)
			return send
		}
	}
	return nil
}
