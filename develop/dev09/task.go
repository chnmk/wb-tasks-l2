package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"slices"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

*/

var links []string

func Wget() {
	GetFlags()
	getPage(input)
}

func getPage(url string) {
	// Если рекурсивное скачивание отключено.
	if !r {
		file, err := os.Create("output.html")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		r, err := http.Get(url)
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
		return
	}

	// Отправляет запрос.
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		log.Fatal("request denied")
	}

	// Считывает тело ответа.
	body := r.Body

	bodyB, err := io.ReadAll(body)
	if err != nil {
		log.Fatal(err)
	}

	bodyS := string(bodyB)

	// Упрощенный поиск ссылок в теле ответа.
	rgx := regexp.MustCompile(`href=".*"`)
	match := rgx.FindAllString(bodyS, -1)

	// Проходит по ссылкам.
	for _, r := range match {
		if !slices.Contains(links, r) {
			links = append(links, r)
			nohref := strings.ReplaceAll(r, "href=", "")
			getPage(url + "/" + strings.ReplaceAll(nohref, "\"", ""))
		}
	}

	createFile(bodyS, url)
}

func createFile(body string, url string) {
	newUrl := strings.ReplaceAll(url, "http://", "")
	newUrl = strings.ReplaceAll(newUrl, "https://", "")
	newUrl = strings.ReplaceAll(newUrl, "//", "")
	newUrl = strings.ReplaceAll(newUrl, "/", "-")
	newUrl = strings.ReplaceAll(newUrl, ".html", "")
	newUrl = strings.TrimRight(newUrl, "-")

	file, err := os.Create(newUrl + ".html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprintln(file, body)
}
