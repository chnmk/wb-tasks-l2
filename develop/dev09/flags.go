package main

import (
	"flag"
	"log"
)

var (
	input string

	r bool // cкачивать рекурсивно
)

func GetFlags() {
	flag.BoolVar(&r, "r", false, "cкачивать рекурсивно")
	flag.Parse()

	if len(flag.Args()) != 1 {
		log.Fatal("ожидалась ссылка")
	}

	input = flag.Arg(0)
}
