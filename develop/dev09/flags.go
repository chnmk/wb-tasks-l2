package main

import (
	"flag"
	"log"
)

var (
	input string

	r string // cкачивать рекурсивно
)

func GetFlags() {
	flag.StringVar(&r, "r", "", "cкачивать рекурсивно")
	flag.Parse()

	if len(flag.Args()) != 1 {
		log.Fatal("ожидалась ссылка")
	}

	input = flag.Arg(0)
}
