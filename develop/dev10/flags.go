package main

import (
	"flag"
	"log"
	"strconv"
	"strings"
)

/*
Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123
*/

var (
	timeout int
	host    string
	port    int
)

func GetFlags() {
	var err error
	var timeout_string string

	flag.StringVar(&timeout_string, "timeout", "3s", "время ожидания в секундах")
	timeout, err = strconv.Atoi(strings.TrimRight(timeout_string, "s"))
	if err != nil {
		log.Fatal(err)
	}

	flag.Parse()

	if len(flag.Args()) != 2 {
		log.Fatal("пример использования: go run . --timeout=3s 1.1.1.1 123")
	}

	host = flag.Arg(0)
	port, err = strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}
}
