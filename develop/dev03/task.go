package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
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
*/

var columnId = 0
var months map[string]int

// Пример использования:
// go run . -r examples/example.txt output.txt
//
// Приоритет флагов при попытке одноврменно использовать противоречащие:
// h > n > M
func main() {
	GetFlags()

	// Чтение файла.
	list, err := read()
	if err != nil {
		log.Fatal(err)
	}

	// Сортировка.
	list, err = sortList(list)
	if err != nil {
		log.Fatal(err)
	}

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

func sortList(list []string) ([]string, error) {
	// Создаёт мапу для сортировки по месяцу.
	if M {
		months = make(map[string]int)
		months["январь"] = 1
		months["февраль"] = 2
		months["март"] = 3
		months["апрель"] = 4
		months["май"] = 5
		months["июнь"] = 6
		months["июль"] = 7
		months["август"] = 8
		months["сентябрь"] = 9
		months["октябрь"] = 10
		months["ноябрь"] = 11
		months["декабрь"] = 12
	}

	// Проверяет, существует ли колонка k для сортировки.
	if k != "" {
		var exists bool
		header := strings.Split(list[0], " ")
		for i, v := range header {
			if v == k {
				exists = true
				columnId = i
				break
			}
		}

		if !exists {
			return []string{}, errors.New("колонка для сортировки не найдена")
		}
	}

	// Проверяет, не отсортированы ли данные.
	if c {
		ok := slices.IsSortedFunc(list, sortFunc)
		if ok {
			log.Println("данные уже отсортированы")
			return list, nil
		}
	}

	// Сортирует данные.
	slices.SortFunc(list, sortFunc)

	// Обратный порядок.
	if r {
		for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
			list[i], list[j] = list[j], list[i]
		}
	}

	// Удаление дубликатов.
	if u {
		list = slices.Compact(list)
	}

	return list, nil
}

func sortFunc(a1, a2 string) int {
	if k != "" {
		a1List := strings.Split(a1, " ")
		a2List := strings.Split(a2, " ")

		if len(a1List) <= columnId || len(a2List) <= columnId {
			log.Print("ошибка: не найдена колонка для строки")
			return 0
		}

		a1 = a1List[columnId]
		a2 = a2List[columnId]
	}

	if h {
		return compareSuffix(a1, a2)
	} else if n {
		return compareNum(a1, a2)
	} else if M {
		return compareMonth(a1, a2)
	}

	if b {
		return strings.Compare(strings.TrimSpace(a1), strings.TrimSpace(a2))
	}

	return strings.Compare(a1, a2)
}

// -n — сортировать по числовому значению;
func compareNum(a, b string) int {
	aInt, err := strconv.Atoi(a)
	if err != nil {
		log.Print("ошибка: строка в колонке чисел")
		return 0
	}

	bInt, err := strconv.Atoi(b)
	if err != nil {
		log.Print("ошибка: строка в колонке чисел")
		return 0
	}

	if aInt > bInt {
		return 1
	} else if aInt < bInt {
		return -1
	}

	return 0
}

// -h — сортировать по числовому значению с учетом суффиксов.
func compareSuffix(a, b string) int {
	aArr := strings.Split(a, ".")
	bArr := strings.Split(b, ".")

	aInt, err := strconv.Atoi(aArr[0])
	if err != nil {
		log.Print("ошибка: строка в колонке чисел")
		return 0
	}

	bInt, err := strconv.Atoi(bArr[0])
	if err != nil {
		log.Print("ошибка: строка в колонке чисел")
		return 0
	}

	if aInt > bInt {
		return 1
	}

	if aInt < bInt {
		return -1
	}

	if len(aArr) > 1 && len(bArr) > 1 {
		aSuf, err := strconv.Atoi(aArr[1])
		if err != nil {
			log.Print("ошибка: строка в колонке чисел")
			return 0
		}

		bSuf, err := strconv.Atoi(bArr[1])
		if err != nil {
			log.Print("ошибка: строка в колонке чисел")
			return 0
		}

		if aSuf > bSuf {
			return 1
		}

		if aSuf < bSuf {
			return -1
		}
	}

	return 0
}

// -M — сортировать по названию месяца;
func compareMonth(a, b string) int {
	if months[strings.ToLower(a)] > months[strings.ToLower(b)] {
		return 1
	} else if months[strings.ToLower(a)] < months[strings.ToLower(b)] {
		return -1
	}

	return 0
}
