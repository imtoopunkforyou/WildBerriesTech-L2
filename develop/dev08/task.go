package main

import (
	"bufio"
	"fmt"
	"github.com/shirou/gopsutil/process"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/*
=== Взаимодействие с ОС ===

# Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	for {
		wd := strings.Split(pwd(), "/")     // Для вывода конкретной директории разбиваем весь путь на элементы слайса.
		fmt.Print(">/", wd[len(wd)-1], " ") // Выводим конкретную директорию, в которой мы находимся

		reader := bufio.NewScanner(os.Stdin) // Создаем reader из stdin.
		if reader.Scan() {                   // Сканируем полученные данные из stdin.
			handleLinuxPipes(reader.Text())
		}
	}
}

func handleLinuxPipes(input string) {
	if len(input) > 0 {
		commands := strings.Split(input, "|") // Разбиваем инпут на пайпы.

		var prevReader io.Reader = nil // Начальный поток ввода.
		for i, command := range commands {
			if strings.TrimSpace(command) != "" { // Проверяем не пустая ли команда в пайпе.
				commandSlice := strings.Fields(command)

				if i == len(commands)-1 { // Проверяем последняя/единственная ли команда.
					execution(commandSlice, prevReader, os.Stdout)
				} else {
					pr, pw := io.Pipe() // Создаем пайп для передачи данных между командами.
					go func(cmd []string, in io.Reader, out io.Writer) {
						execution(cmd, in, out)
						pw.Close() // Закрываем writer после выполнения команды.
					}(commandSlice, prevReader, pw)

					prevReader = pr // Обновляем поток ввода для следующей команды
				}
			}
		}
	}
}

// Вызов команд для терминала.
func execution(str []string, r io.Reader, w io.Writer) {
	switch str[0] { // Смотрим какая команда пришла.
	case "pwd":
		fmt.Fprintln(w, pwd())
	case "echo":
		fmt.Fprintln(w, echo(str))
	case "kill":
		kill(str)
	case "ps":
		ps(w)
	case "cd":
		cd(str)
	case "\\exit":
		os.Exit(0)
	default:
		forkExec(str, r, w)
	}
}

func forkExec(str []string, r io.Reader, w io.Writer) {
	cmd := exec.Command(str[0], str[1:]...) // Вызываем команды из терминала.
	cmd.Stdin = r                           // Назначаем вводом reader
	cmd.Stdout = w                          // Назначаем выводом writer
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintln(w, err)
	}
}

func echo(str []string) string {
	var res strings.Builder
	for _, line := range str[1:] {
		line = strings.Trim(line, "\"")
		line = strings.Trim(line, "'")
		res.WriteString(line + " ")
	}
	return res.String()
}

func pwd() string {
	wd, err := os.Getwd()
	if err != nil {
		return err.Error()
	}
	return wd
}

func kill(str []string) {
	if len(str) < 2 {
		fmt.Println("kill: not enough arguments")
		return
	}

	for _, str := range str[1:] {
		id, err := strconv.Atoi(str) // Получаем pid из строки.
		if err != nil {
			fmt.Println(err)
			continue // Если ошибка, то идем дальше по слайсу.
		}

		proc, err := os.FindProcess(id) // Ищем процесс, если получили его из строки.
		if err != nil {
			fmt.Println(err)
			continue // Если ошибка, то идем дальше по слайсу.
		}

		err = proc.Kill()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func ps(w io.Writer) {
	proc, err := process.Processes() // Получаем процессы.
	if err != nil {
		fmt.Fprintln(w, err)
	}

	fmt.Fprintf(w, "PID\t TIME\t CMD\n")
	for _, p := range proc {
		time, _ := p.Times()  // Получаем время работы процесса.
		cmd, _ := p.Cmdline() // Получаем имя исполнителя процесса.
		fmt.Fprintf(w, "%v\t %0.2f\t %v\n", p.Pid, time.Total(), cmd)
	}
}

func cd(str []string) {
	if len(str) == 1 || str[1] == "~" {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
		} else { // Если у нас получилось взять домашнюю директорию, то переходим в нее.
			if err = os.Chdir(home); err != nil {
				fmt.Println(err)
			}
		}
	} else {
		if err := os.Chdir(str[1]); err != nil {
			fmt.Println(err)
		}
	}
}
