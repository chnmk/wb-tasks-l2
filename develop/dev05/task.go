package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
*/

// Пример использования:
// go run . -A 1 examples/example.txt output.txt Втор
func main() {
	GetFlags()

	// Приоритет флагов при попытке одноврменно использовать противоречащие:
	// C > A > B
	if B != 0 && A != 0 {
		B = 0
		if C != 0 {
			B = 0
		}
	}

	// Чтение файла.
	list, err := read()
	if err != nil {
		log.Fatal(err)
	}

	// Фильтрация.
	list = filterFile(list)

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

func filterFile(list []string) []string {
	var result []string

	// Ограничить количество строк.
	if c != 0 {
		list = list[:c]
	}

	var found []int
	for num, s := range list {
		// Игнорировать регистр.
		if i {
			s = strings.ToLower(s)
		}

		// Точное совпадение со строкой, не паттерн.
		if F {
			if s == pattern {
				found = append(found, num)
			}
			continue
		}

		// Если нужен паттерн.
		if strings.Contains(s, pattern) {
			found = append(found, num)
		}
	}

	for _, num := range found {
		// Печатать +N строк после совпадения.
		if A != 0 {
			result = append(result, list[num])

			for add := 1; add <= A; add++ {
				if len(list) > num+add {
					new := list[num+add]
					if n {
						new = strconv.Itoa(num+1) + " " + new
					}

					result = append(result, new)
				}
			}
			continue
		}

		// Печатать +N строк до совпадения.
		if B != 0 {
			for add := B; add > 0; add-- {
				if num-add-1 >= 0 {
					new := list[num-add]
					if n {
						new = strconv.Itoa(num+1) + " " + new
					}

					result = append(result, new)
				}
			}

			result = append(result, list[num])
			continue
		}

		// Печатать ±N строк вокруг совпадения.
		if C != 0 {
			for add := C; add > 0; add-- {
				if num-add-1 >= 0 {
					new := list[num-add]
					if n {
						new = strconv.Itoa(num+1) + " " + new
					}

					result = append(result, new)
				}
			}

			result = append(result, list[num])

			for add := 1; add <= C; add++ {
				if len(list) > num+add {
					new := list[num+add]
					if n {
						new = strconv.Itoa(num+1) + " " + new
					}

					result = append(result, new)
				}
			}
			continue
		}

		new := list[num]
		if n {
			new = strconv.Itoa(num+1) + " " + new
		}

		result = append(result, new)
	}

	// Вместо совпадения, исключать.
	if v {
		var result2 []string

		j := 0
		for num, s := range list {
			if len(result) > j {
				if s == result[j] {
					j++
					continue
				}
			}

			if n {
				s = strconv.Itoa(num+1) + " " + s
			}
			result2 = append(result2, s)
		}

		return result2
	}

	return result
}
