package main

import (
	"fmt"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func init() {
	GetFlags()
}

func main() {
	// Получить STDIN (в бесконечном цикле, до сигнала о завершении, и сразу выводить?)
	// header := "Example string 123"
	string1 := "Exampl abc 1"
	string2 := "Examp bca 2"
	string3 := "Exam cba 3"

	new1 := strings.Split(string1, " ")
	fmt.Println(new1[0])

	new2 := strings.Split(string2, " ")
	fmt.Println(new2[0])

	new3 := strings.Split(string3, " ")
	fmt.Println(new3[0])
}
