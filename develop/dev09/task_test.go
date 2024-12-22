package main

import (
	"bufio"
	"net/http"
	"os"
	"strings"
	"sync"
	"testing"
	"time"
)

var once sync.Once

func TestWgetDefault(t *testing.T) {
	once.Do(func() {
		http.Handle("/", http.FileServer(http.Dir("./example")))
		go http.ListenAndServe(":3000", nil)

		time.Sleep(1 * time.Second)
	})

	getPage("http://127.0.0.1:3000/page1.html")

	p1, err := os.Open("output.html")
	if err != nil {
		t.Fatal("file not found")
	}

	defer p1.Close()

	var found bool

	scanner := bufio.NewScanner(p1)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "Страница 1") {
			found = true
		}
	}
	if err := scanner.Err(); err != nil {
		t.Fatal("error reading file")
	}

	if !found {
		t.Error("wrong result")
	}

	os.Remove("output.html")
}

func TestWgetRecursive(t *testing.T) {
	r = true
	t.Cleanup(func() {
		r = false
	})

	once.Do(func() {
		http.Handle("/", http.FileServer(http.Dir("./example")))
		go http.ListenAndServe(":3000", nil)

		time.Sleep(1 * time.Second)
	})

	getPage("http://127.0.0.1:3000/")

	if len(links) != 3 {
		t.Error("wrong number of pages")
	}

	p1, err := os.Open("127.0.0.1:3000page1.html")
	if err != nil {
		t.Fatal("file not found")
	}

	defer p1.Close()

	var found bool

	scanner := bufio.NewScanner(p1)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "Страница 1") {
			found = true
		}
	}
	if err := scanner.Err(); err != nil {
		t.Fatal("error reading file")
	}

	if !found {
		t.Error("wrong result")
	}

	os.Remove("127.0.0.1:3000.html")
	os.Remove("127.0.0.1:3000page1.html")
	os.Remove("127.0.0.1:3000page2.html")
}
