package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

- cd <args> — смена директории (в качестве аргумента могут быть то-то и то);

- pwd — показать путь до текущего каталога;

- echo <args> — вывод аргумента в STDOUT;

- kill <args> — «убить» процесс, переданный в качесте аргумента (пример: такой-то пример);

- ps — выводит общую информацию по запущенным процессам в формате такой-то формат.

Так же требуется поддерживать функционал fork/exec-команд.

Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

*Шелл — это обычная консольная программа, которая будучи запущенной,
в интерактивном сеансе выводит некое приглашение в STDOUT и ожидает ввода пользователя через STDIN.
Дождавшись ввода, обрабатывает команду согласно своей логике и при необходимости выводит результат на экран.
Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).
*/

func main() {
	r := bufio.NewReader(os.Stdin)

	for {
		input, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}

		command := strings.Fields(strings.TrimSuffix(input, "\n"))

		if len(command) == 0 {
			fmt.Println("invalid input")
			continue
		}

		switch command[0] {
		case "quit":
			os.Exit(0)
		case "cd":
			cd(command)
		}
	}
}

func cd(command []string) {
	if len(command) != 2 {
		fmt.Println("invalid input")
		return
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	err = os.Chdir(homeDir)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
}
