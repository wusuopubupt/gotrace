package main

import (
	"container/list"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/divan/gotrace/trace"
)

func ConvertEvents(events []*trace.Event) (Commands, error) {
	var c Commands

	sends := list.New()

	debug := os.Getenv("GOTRACE_DEBUG") == "1"
	var lastG uint64
	for _, ev := range events {
		switch ev.Type {
		case trace.EvGoStart:
			fmt.Println(" ---> GoStart:", ev.G, "from", lastG)
			if ev.G == 1 && lastG == 0 {
				c.StartGoroutine(ev.Ts, ev.Stk[0].Fn, ev.G, lastG)
			}
			lastG = ev.G
		case trace.EvGoCreate:
			if len(ev.Stk) > 0 {
				if strings.HasPrefix(ev.Stk[0].Fn, "runtime") {
					if ev.Stk[0].Fn != "runtime.main" {
						break
					}
				}
				c.StartGoroutine(ev.Ts, ev.Stk[0].Fn, ev.Args[0], ev.G)
				if debug {
					fmt.Println(" ---> Create:", ev.Args[0], "from", ev.G)
				}
			}
		case trace.EvGoUnblock:
			lastG = ev.Args[0]
			fmt.Println("Go: Unblock: set lastG to", lastG)
		case trace.EvGoEnd:
			c.StopGoroutine(ev.Ts, "", ev.G)
			lastG = ev.G
			if debug {
				fmt.Println("End:", ev.G)
			}
		case trace.EvGoSend:
			sends.PushBack(ev)
			if debug {
				fmt.Printf("[DD] %d, Send: G:%d, CH: %d, EvID: %d, Val:%d\n", ev.Ts, ev.G, ev.Args[1], ev.Args[0], ev.Args[2])
			}
		case trace.EvGoRecv:
			if debug {
				fmt.Printf("[DD] %d, Recv: G:%d, CH: %d, EvID: %d, Val:%d - %v\n", ev.Ts, ev.G, ev.Args[1], ev.Args[0], ev.Args[2], ev)
			}
			send := findSource(sends, ev)
			if send == nil {
				// it's either channel close() or error in findSource
				continue
			}
			if debug {
				fmt.Printf("[DD] %d, Recv->Send: FromG:%d, ToG: %d, CH: %d, EvID: %d, Val:%d\n", send.Ts, send.G, ev.G, ev.Args[1], ev.Args[0], ev.Args[2])
			}
			c.ChanSend(send.Ts, ev.Args[1], send.G, ev.G, send.Args[2])
		case trace.EvGCStart, trace.EvGCDone, trace.EvGCScanStart, trace.EvGCScanDone:
			lastG = 1
			/*
				case trace.EvGoSched, trace.EvGoPreempt,
					trace.EvGoSleep, trace.EvGoBlock, trace.EvGoBlockSelect, trace.EvGoBlockSend, trace.EvGoBlockRecv,
					trace.EvGoBlockSync, trace.EvGoBlockCond, trace.EvGoBlockNet,
					trace.EvGoSysBlock:
					fmt.Println("Ev:", ev.Type, ev.G, ev.Args)
					//lastG = ev.G
			*/
		case trace.EvGoStop:
			if debug {
				fmt.Println("Stop:", ev.G)
			}
			lastG = 1
		default:
			if debug {
				fmt.Println("Ev:", ev.Type, ev.G, ev.Args)
			}
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
