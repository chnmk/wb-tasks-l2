package main

import (
	"flag"
	"log"
)

/*
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

var (
	input  string // Файл .txt, который необходимо отсортировать.
	output string // Файл .txt, в который записывается результат.

	k string // Указание колонки для сортировки.
	n bool   // Cортировать по числовому значению.
	r bool   // Cортировать в обратном порядке.
	u bool   // Не выводить повторяющиеся строки.

	M bool // Сортировать по названию месяца.
	b bool // Игнорировать хвостовые пробелы.
	c bool // Проверять отсортированы ли данные.
	h bool // Сортировать по числовому значению с учетом суффиксов
)

func GetFlags() {
	flag.StringVar(&k, "k", "", "указание колонки для сортировки")
	flag.BoolVar(&n, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&r, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&u, "u", false, "не выводить повторяющиеся строки")
	flag.BoolVar(&M, "M", false, "сортировать по названию месяца")
	flag.BoolVar(&b, "b", false, "игнорировать хвостовые пробелы")
	flag.BoolVar(&c, "c", false, "проверять отсортированы ли данные")
	flag.BoolVar(&h, "h", false, "сортировать по числовому значению с учетом суффиксов")

	flag.Parse()

	if len(flag.Args()) != 2 {
		log.Fatal("ожидались файлы для чтения и записи")
	}

	input = flag.Arg(0)
	output = flag.Arg(1)
	if input[len(input)-4:] != ".txt" || output[len(output)-4:] != ".txt" {
		log.Fatal("ожидались файлы для чтения и записи")
	}
}
