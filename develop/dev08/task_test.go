package main

import (
	"bytes"
	"log"
	"os"
	"testing"
)

func TestCmdDefault(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	log.SetFlags(0)

	t.Cleanup(func() { log.SetOutput(os.Stderr) })

	selectCommand([]string{"echo", "test1"})

	if buf.String() != "test1\n" {
		t.Errorf("wrong value, got: %s", buf.String())
	}
}

func TestCmdPipe(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	log.SetFlags(0)

	t.Cleanup(func() { log.SetOutput(os.Stderr) })

	selectCommand([]string{"echo", "test1", "|", "echo", "test2"})

	if buf.String() != "test1\ntest2\n" {
		t.Errorf("wrong value, got: %s", buf.String())
	}
}
