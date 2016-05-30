package main

import (
	"encoding/json"
	"fmt"
)

type Commands []*Command

// Command is a common structure for all
// types of supported events (aka 'commands').
// It's main purpose to handle JSON marshalling.
type Command struct {
	Time     int64       "json:\"t\""
	Command  string      "json:\"command\""
	Name     string      "json:\"name,omitempty\""
	Parent   string      "json:\"parent,omitempty\""
	Channels []string    "json:\"channels,omitempty\""
	From     string      "json:\"from,omitempty\""
	To       string      "json:\"to,omitempty\""
	Channel  string      "json:\"ch,omitempty\""
	Value    interface{} "json:\"value,omitempty\""
	EventID  string      "json:\"eid,omitempty\""
	Duration int64       "json:\"duration,omitempty\""
}

func (c *Commands) toJSON() []byte {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		panic(err)
	}

	return data
}

func (c *Commands) StartGoroutine(ts int64, name string, gid, pid uint64) {
	parent := fmt.Sprintf("#%d", pid)
	// ignore parent for 'main()' which has pid 0
	if pid == 0 {
		parent = ""
	}
	cmd := &Command{
		Time:    ts,
		Command: "create goroutine",
		Name:    fmt.Sprintf("#%d", gid),
		Parent:  parent,
	}
	*c = append(*c, cmd)
}

func (c *Commands) StopGoroutine(ts int64, name string, gid uint64) {
	cmd := &Command{
		Time:    ts,
		Command: "stop goroutine",
		Name:    fmt.Sprintf("#%d", gid),
	}
	*c = append(*c, cmd)
}

func (c *Commands) ChanSend(ts int64, cid, gid, eid, val uint64) {
	cmd := &Command{
		Time:    ts,
		Command: "start send",
		From:    fmt.Sprintf("#%d", gid),
		Channel: fmt.Sprintf("#%d", cid),
		Value:   fmt.Sprintf("%d", val),
		EventID: fmt.Sprintf("%d", eid),
	}
	*c = append(*c, cmd)
}

func (c *Commands) ChanRecv(ts int64, cid, gid, eid, val uint64) {
	cmd := &Command{
		Time:    ts,
		Command: "start recv",
		To:      fmt.Sprintf("#%d", gid),
		Channel: fmt.Sprintf("#%d", cid),
		Value:   fmt.Sprintf("%d", val),
		EventID: fmt.Sprintf("%d", eid),
	}
	*c = append(*c, cmd)
}
