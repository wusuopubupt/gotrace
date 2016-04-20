package main

import (
	"encoding/json"
	"fmt"
	"io"
)

// cmdWriter writes serialized commands
// into the given writer.
type cmdWriter struct {
	w io.Writer
}

// NewCmdWriter inits new cmdWriter.
func NewCmdWriter(w io.Writer) *cmdWriter {
	return &cmdWriter{
		w: w,
	}
}

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
	Duration int64       "json:\"duration,omitempty\""
}

func (c *cmdWriter) write(cmd *Command) {
	data, err := json.Marshal(cmd)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(c.w, string(data))
}

func (c *cmdWriter) StartGoroutine(ts int64, name string, gid, pid uint64) {
	cmd := &Command{
		Time:    ts,
		Command: "create goroutine",
		Name:    fmt.Sprintf("#%d", gid),
		Parent:  fmt.Sprintf("#%d", pid),
	}
	c.write(cmd)
}

func (c *cmdWriter) StopGoroutine(ts int64, name string, gid uint64) {
	cmd := &Command{
		Time:    ts,
		Command: "stop goroutine",
		Name:    fmt.Sprintf("#%d", gid),
	}
	c.write(cmd)
}

func (c *cmdWriter) ChanSend(ts int64, cid, gid, did uint64) {
	cmd := &Command{
		Time:    ts,
		Command: "start send",
		From:    fmt.Sprintf("#%d", gid),
		Channel: fmt.Sprintf("#%d", cid),
		Value:   fmt.Sprintf("%d%d", cid, did),
	}
	c.write(cmd)
}

func (c *cmdWriter) ChanRecv(ts int64, cid, gid, did uint64) {
	cmd := &Command{
		Time:    ts,
		Command: "start recv",
		To:      fmt.Sprintf("#%d", gid),
		Channel: fmt.Sprintf("#%d", cid),
		Value:   fmt.Sprintf("%d%d", cid, did),
	}
	c.write(cmd)
}
