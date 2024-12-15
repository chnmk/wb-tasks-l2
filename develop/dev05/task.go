package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func init() {
	GetFlags()
}

func main() {
	file, err := os.Open("examples/example.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scan := bufio.NewScanner(file)

	for scan.Scan() {
		text := scan.Text()

		// TODO: получать необходимый паттерн при запуске утилиты
		// TODO: условия по флагам
		if strings.Contains(text, "ы") {
			fmt.Println(text)
		}
	}
}
