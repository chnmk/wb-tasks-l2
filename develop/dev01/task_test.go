package CurrentTime

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func TestGetTime(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	t.Cleanup(func() { log.SetOutput(os.Stderr) }) // Возврат к значению по умолчанию

	GetCurrentTime()

	_, err := time.Parse(time.TimeOnly, strings.Split(buf.String(), " ")[1])
	if err != nil {
		t.Errorf("error: expected valid time format: %v", err)
	}
}

func TestGetTimeTwice(t *testing.T) {
	var buf1 bytes.Buffer
	log.SetOutput(&buf1)

	t.Cleanup(func() { log.SetOutput(os.Stderr) }) // Возврат к значению по умолчанию

	GetCurrentTime()

	time.Sleep(1 * time.Second)

	var buf2 bytes.Buffer
	log.SetOutput(&buf2)

	GetCurrentTime()

	t1, err := time.Parse(time.TimeOnly, strings.Split(buf1.String(), " ")[1])
	if err != nil {
		t.Fatal(err)
	}

	t2, err := time.Parse(time.TimeOnly, strings.Split(buf2.String(), " ")[1])
	if err != nil {
		t.Fatal(err)
	}

	if t2.Before(t1) {
		t.Errorf("time 2 is before time 1")
	}
}
