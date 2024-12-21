package main

import (
	"testing"
	"time"
)

func TestChanDefault(t *testing.T) {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Second),
		sig(5*time.Second),
		sig(6*time.Second),
	)

	if time.Since(start) > 3*time.Second || time.Since(start) < 1*time.Second {
		t.Error("invalid duration")
	}
}
