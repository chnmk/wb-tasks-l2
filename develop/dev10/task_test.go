package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestTelnetTimeout(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	timeout = 1
	host = "googlejenfwehbfiwebfieb.com"
	port = 9999

	t.Cleanup(func() {
		log.SetOutput(os.Stderr)
		timeout = 3
		host = ""
		port = 0
	})

	connect()

	if !strings.Contains(buf.String(), "i/o timeout") {
		t.Errorf("wrong result, got: %s", buf.String())
	}
}

func TestTelnetDefault(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	timeout = 1
	host = "localhost"
	port = 3000

	t.Cleanup(func() {
		log.SetOutput(os.Stderr)
		timeout = 3
		host = ""
		port = 0
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test"))
	})
	go http.ListenAndServe(":3000", nil)

	connect()

	if !strings.Contains(buf.String(), "shutting down...") {
		t.Errorf("wrong result, got: %s", buf.String())
	}
}
