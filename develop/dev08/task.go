package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
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
