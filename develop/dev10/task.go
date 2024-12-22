package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	GetFlags()

	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), time.Duration(timeout)*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	r := bufio.NewReader(os.Stdin)

	for {
		// Отправляет сообщение.
		input, err := r.ReadString('\n')
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Fprintln(conn, input)

		// Печатает полученные данные.
		status, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println(status)
	}
}
