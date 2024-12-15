package main

import "flag"

/*
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
*/

var (
	f string
	d string
	s bool
)

func GetFlags() {
	flag.StringVar(&f, "k", "", "TODO")
	flag.StringVar(&d, "n", "", "TODO")
	flag.BoolVar(&s, "r", false, "TODO")

	flag.Parse()
}
