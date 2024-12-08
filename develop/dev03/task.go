package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

/*
=== Утилита sort ===

Отсортировать строки в файле по аналогии с консольной утилитой sort (man sort — смотрим описание и основные параметры):
на входе подается файл из несортированными строками, на выходе — файл с отсортированными.


Реализовать поддержку утилитой следующих ключей:

-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел);

-n — сортировать по числовому значению;

-r — сортировать в обратном порядке;

-u — не выводить повторяющиеся строки.

Дополнительно
Реализовать поддержку утилитой следующих ключей:

-M — сортировать по названию месяца;

-b — игнорировать хвостовые пробелы;

-c — проверять отсортированы ли данные;

-h — сортировать по числовому значению с учетом суффиксов.


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

	var list []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Подумать
		list = append(list, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Strings(list)

	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}
}
