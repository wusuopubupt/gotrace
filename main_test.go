package main

import (
	"path/filepath"
	"testing"
)

func TestExamples(t *testing.T) {
	tests := []struct {
		path     string
		cmdCount int
	}{
		{"server1.go", 5},
		{"workers1.go", 126},
		{"workers2.go", 346},
		{"fanin1.go", 48},
		{"hello01.go", 5},
		{"pingpong01.go", 16},
		{"pingpong02.go", 18},
		{"pingpong03.go", 212},
		{"primesieve.go", 188},
		{"select.go", 6},
	}

	for _, test := range tests {
		path := filepath.Join("examples", test.path)
		t.Log("Testing", path)
		src := NewNativeRun(path)
		events, err := src.Events()
		if err != nil {
			t.Fatal(path, err)
		}
		commands, err := ConvertEvents(events)
		if err != nil {
			t.Fatal(path, err)
		}

		count := len(commands)
		if count != test.cmdCount {
			t.Fatalf("Wrong number of commands: %s: %d", path, count)
		}
	}
}
