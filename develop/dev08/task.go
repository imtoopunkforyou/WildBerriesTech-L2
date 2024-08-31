package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/*
Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

cd - смена директории (в качестве аргумента могут быть то-то и то)
pwd - показать путь до текущего каталога
echo - вывод аргумента в STDOUT
kill - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
ps - выводит общую информацию по запущенным процессам в формате такой-то формат
Так же требуется поддерживать функционал fork/exec-команд

Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).
*/
func main() {
	var input string

	for {
		fmt.Print(">")
		fmt.Scanln(&input)
		New(strings.Fields(input), input)
	}
}

func New(strSlice []string, str string) {
	switch strSlice[0] {
	case "pwd":
		fmt.Println(pwd())
	case "echo":
		fmt.Println(strings.Trim(str, "echo \""))
	case "kill":
		kill(strSlice)
	case "ps":
		ps()
	case "cd":
		cd()
	case "\\exit":
		os.Exit(0)
	default:
		fmt.Println(forkExec(strSlice))
	}
}

func pwd() string {
	wd, err := os.Getwd()
	if err != nil {
		return err.Error()
	}
	return wd
}

func forkExec(strSlice []string) string {
	cmd := exec.Command(strSlice[0], strSlice[1:]...)
	res, err := cmd.CombinedOutput()
	if err != nil {
		return err.Error()
	}

	if len(res) > 0 {
		return string(res[:len(res)-1]) // Избавляемся от последней пустой слайса.
	}
	return ""
}

func kill(strSlice []string) {
	for _, str := range strSlice[1:] {
		id, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println(err)
			continue
		}

		proc, err := os.FindProcess(id)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = proc.Kill()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func ps() {
	fmt.Println("kakish")
}

func cd() {
	fmt.Println("aboba")
}
