package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
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

var currentDir string

func main() {
	executable, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	currentDir = filepath.Dir(executable)

	log.SetFlags(0)

	r := bufio.NewReader(os.Stdin)

	for {
		input, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}

		command := strings.Fields(strings.TrimSuffix(input, "\n"))

		if len(command) == 0 {
			log.Println("invalid input")
			continue
		}

		selectCommand(command)
	}
}

func selectCommand(input []string) {
	str := strings.Join(input, " ")
	if strings.Contains(str, "|") {
		newStr := strings.Split(str, "|")
		for _, cmd := range newStr {
			cmd = strings.TrimLeft(cmd, " ")
			cmd = strings.TrimRight(cmd, " ")
			selectCommand(strings.Split(cmd, " "))
		}

		return
	}

	switch input[0] {
	case "quit":
		os.Exit(0)
	case "cd":
		cd(input)
	case "pwd":
		pwd()
	case "echo":
		echo(input)
	case "kill":
		kill(input)
	case "ps":
		ps()
	case "fork":
		fork(input)
	case "exec":
		exe(input)
	default:
		log.Println("invalid input")
	}
}

func cd(command []string) {
	if len(command) != 2 {
		log.Println("invalid input")
		return
	}

	if command[1] == "./.." {
		temp := strings.Split(currentDir, "/")
		temp = temp[:len(temp)-1]
		currentDir = strings.Join(temp, "/")

		err := os.Chdir(currentDir)
		if err != nil {
			log.Printf("error: %v", err)
			return
		}

		return
	}

	currentDir = currentDir + "/" + command[1]

	err := os.Chdir(currentDir)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
}

func pwd() {
	log.Println(filepath.Dir(currentDir) + "/" + filepath.Base(currentDir))
}

func echo(command []string) {
	if len(command) < 2 {
		log.Println("invalid input")
		return
	}

	for i, text := range command {
		if i != 0 {
			log.Println(text)
		}
	}
}

func kill(command []string) {
	if len(command) != 2 {
		log.Println("invalid input")
		return
	}

	id, err := strconv.Atoi(command[1])
	if err != nil {
		log.Println("invalid input")
		return
	}

	process, err := os.FindProcess(id)
	if err != nil {
		log.Println(err)
	}

	process.Kill()
}

func ps() {
	processFiles, err := filepath.Glob("/proc/*/exe")
	if err != nil {
		log.Println("internal error")
		return
	}

	for _, file := range processFiles {
		id := strings.Split(file, "/")[2]
		target, _ := os.Readlink(file)
		if len(target) > 0 {
			log.Printf("%s %s\n", id, target)
		}
	}
}

func fork(command []string) {
	if len(command) < 2 {
		log.Println("invalid input")
		return
	}

	pid, err := syscall.ForkExec(command[1], command[2:], nil)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("process started, pid\n: %d", pid)
}

func exe(command []string) {
	if len(command) < 2 {
		log.Println("invalid input")
		return
	}

	args := strings.Join(command[1:], " ")

	cmd := exec.Command(args)
	if err := cmd.Run(); err != nil {
		log.Println(err)
	}
}
