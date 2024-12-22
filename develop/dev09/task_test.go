package main

import (
	"log"
	"net/http"
	"testing"
	"time"
)

func TestWgetRecursive(t *testing.T) {
	http.Handle("/", http.FileServer(http.Dir("./example")))
	go http.ListenAndServe(":3000", nil)

	time.Sleep(1 * time.Second)

	getPage("http://127.0.0.1:3000/")
	log.Println(links)

	/*
		if status := responseRecorder.Code; status != http.StatusBadRequest {
			t.Errorf("expected status code: %d, got %d", http.StatusBadRequest, status)
		}
	*/
}
