package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	file, err := os.Create("example.html")
	if err != nil {
		panic(
			err)
	}
	defer file.Close()

	r, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		log.Fatal("request denied")
	}

	_, err = io.Copy(file, r.Body)
	if err != nil {
		panic(err)
	}
}
