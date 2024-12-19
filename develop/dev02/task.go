package main

import (
	"log"
	"regexp"
	"strconv"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func Unpack(input string) string {
	// Проверка на пустую строку.
	if input == "" {
		return ""
	}

	// Проверка на наличие не только цифр.
	ok, err := regexp.MatchString("[^0-9]", input)
	if err != nil {
		log.Print(err)
		return ""
	}
	if !ok {
		log.Print("некорректная строка")
		return ""
	}

	// Проверка на наличие цифр.
	ok, err = regexp.MatchString("[0-9]", input)
	if err != nil {
		log.Print(err)
		return ""
	}
	if !ok {
		return input
	}

	var result []rune
	var prev rune
	var preprev rune
	for _, r := range input {

		// Помечает последовательности вида "\\5".
		if string(prev) == "\\" && string(preprev) != "\\" {
			preprev = prev
			prev = r
			result = append(result, r)
			continue
		}
		preprev = 0

		// Экспейп-символ пропускается.
		if string(r) == "\\" {
			prev = r
			continue
		}

		// Если не число - просто добавляет символ.
		if !unicode.IsDigit(r) {
			prev = r
			result = append(result, r)
			continue
		}

		// Если число - добавляет предыдущий символ до нужного количества.
		num, err := strconv.Atoi(string(r))
		if err != nil {
			log.Print("внутренняя ошибка")
			return ""
		}
		for i := 0; i < num-1; i++ {
			result = append(result, prev)
		}

	}

	return string(result)
}
