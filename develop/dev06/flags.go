package main

import (
	"flag"
	"log"
	"strings"
)

/*
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
*/

var (
	input  string // Файл .txt, который необходимо разделить.
	output string // Файл .txt, в который записывается результат.

	f string // выбрать поля (колонки)
	d string // использовать другой разделитель
	s bool   // только строки с разделителем

	f_split []string
)

func GetFlags() {
	flag.StringVar(&f, "f", "", "выбрать поля (колонки)")
	flag.StringVar(&d, "d", "", "использовать другой разделитель")
	flag.BoolVar(&s, "s", false, "только строки с разделителем")

	flag.Parse()

	f_split = strings.Split(f, " ")

	if len(flag.Args()) != 2 {
		log.Fatal("ожидались файлы для чтения и записи")
	}

	input = flag.Arg(0)
	output = flag.Arg(1)
	if input[len(input)-4:] != ".txt" || output[len(output)-4:] != ".txt" {
		log.Fatal("ожидались файлы для чтения и записи")
	}
}
