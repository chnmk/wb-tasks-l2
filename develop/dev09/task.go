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
	// Отправляет запрос.
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		log.Fatal("request denied")
	}

	fmt.Println(r.Request)

	// Считывает тело ответа.
	body := r.Body

	bodyB, err := io.ReadAll(body)
	if err != nil {
		log.Fatal(err)
	}

	bodyS := string(bodyB)

	// Ищет ссылки в теле ответа.
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
	newUrl = strings.ReplaceAll(newUrl, "/", "-")

	file, err := os.Create(newUrl + ".html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprintln(file, body)
}

/*
func wget(url string) {
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		log.Fatal("request denied")
	}

	body := r.Body

	recursive(body)
	createFile(body)

}
*/

/*
func recursive(body io.ReadCloser) {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)

	r := regexp.MustCompile(`href=".*"`)
	matches := r.FindAllString(bodyString, -1)

	for _, r := range matches {
		fmt.Println(r)
		fmt.Println(" ")
	}
}
*/

/*
func main() {
	file, err := os.Create("example.html")
	if err != nil {
		panic(err)
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

	recursive(r.Body)

	_, err = io.Copy(file, r.Body)
	if err != nil {
		panic(err)
	}

}

func recursive(b io.ReadCloser) {
	bodyBytes, err := io.ReadAll(b)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	log.Println(bodyString, "f")
}
*/
