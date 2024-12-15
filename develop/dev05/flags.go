package main

import "flag"

/*
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки
*/

var (
	A int
	B int
	C int
	c int

	i bool
	v bool
	F bool
	n bool
)

func GetFlags() {
	flag.IntVar(&A, "A", 1, "TODO")
	flag.IntVar(&B, "B", 1, "TODO")
	flag.IntVar(&C, "C", 1, "TODO")
	flag.IntVar(&c, "c", 1, "TODO")
	flag.BoolVar(&i, "i", false, "TODO")
	flag.BoolVar(&v, "v", false, "TODO")
	flag.BoolVar(&F, "F", false, "TODO")
	flag.BoolVar(&n, "n", false, "TODO")

	flag.Parse()
}
