package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

/*
=== Утилита cut ===

Реализовать утилиту аналог консольной команды cut (man cut).
Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.

Реализовать поддержку утилитой следующих ключей:
-f — "fields": выбрать поля (колонки);
-d — "delimiter": использовать другой разделитель;
-s — "separated": только строки с разделителем.
*/

// Пример использования:
// go run . -f "Id Name" -d " " examples/example.txt output.txt
func main() {
	GetFlags()

	// Чтение файла.
	list, err := read()
	if err != nil {
		log.Fatal(err)
	}

	// Разделение.
	list = cut(list)

	// Запись результата.
	err = write(output, list)
	if err != nil {
		log.Fatal(err)
	}
}

func read() ([]string, error) {
	var list []string

	file, err := os.Open(input)
	if err != nil {
		return []string{}, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return []string{}, err
	}

	return list, nil
}

func write(path string, text []string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	for _, s := range text {
		fmt.Fprintln(file, s)
	}

	return nil
}

func cut(list []string) []string {
	var result []string
	var columnIds []int

	if d == "" {
		d = "\t"
	}

	// Работа со строкой заголовка.
	columns := strings.Split(list[0], d)
	for id, title := range columns {
		if slices.Contains(f_split, title) {
			columnIds = append(columnIds, id)
		}
	}

	for _, line := range list {
		// Только строки с разделителем.
		if s {
			if !strings.Contains(line, d) {
				continue
			}
		}

		var newLine string
		columns := strings.Split(line, d)
		for id, text := range columns {
			if slices.Contains(columnIds, id) {
				newLine = newLine + d + text
			}
		}

		result = append(result, strings.TrimLeft(newLine, d))
	}

	return result
}
