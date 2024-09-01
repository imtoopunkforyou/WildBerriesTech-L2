package main

import (
	"bufio"
	"fmt"
	"github.com/shirou/gopsutil/process"
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
	for {
		wd := strings.Split(pwd(), "/") // Для вывода конкретной директории разбиваем весь путь на элементы слайса.

		fmt.Print(">/", wd[len(wd)-1], " ")  // Выводим конкретную директорию, в которой мы находимся
		reader := bufio.NewScanner(os.Stdin) // Создаем reader из stdin.
		if reader.Scan() {                   // Сканируем полученные данные из stdin.
			New(strings.Fields(reader.Text()), reader.Text()) // Передача reader в виде строки нужно только для echo.
		}
	}
}

// New Создание нового терминала.
func New(strSlice []string, str string) {
	switch strSlice[0] { // Смотрим какая команда пришла.
	case "pwd":
		fmt.Println(pwd())
	case "echo":
		fmt.Println(strings.Trim(str, "echo \""))
	case "kill":
		kill(strSlice)
	case "ps":
		ps()
	case "cd":
		cd(strSlice)
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
	cmd := exec.Command(strSlice[0], strSlice[1:]...) // Вызываем команды из терминала.
	res, err := cmd.CombinedOutput()                  // Получаем вывод.
	if err != nil {
		return err.Error()
	}

	if len(res) > 0 {
		return string(res[:len(res)-1]) // Избавляемся от последней пустой строки слайса.
	}
	return ""
}

func kill(strSlice []string) {
	for _, str := range strSlice[1:] {
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

func ps() {
	proc, err := process.Processes() // Получаем процессы.
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("PID\t TIME\t CMD\n")
	for _, p := range proc {
		time, _ := p.Times()  // Получаем время работы процесса.
		cmd, _ := p.Cmdline() // Получаем имя исполнителя процесса.
		fmt.Printf("%v\t %0.2f\t %v\n", p.Pid, time.Total(), cmd)
	}
}

func cd(strSlice []string) {
	if len(strSlice) == 1 || strSlice[1] == "~" {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
		} else { // Если у нас получилось взять домашнюю директорию, то переходим в нее.
			if err = os.Chdir(home); err != nil {
				fmt.Println(err)
			}
		}
	} else {
		if err := os.Chdir(strSlice[1]); err != nil {
			fmt.Println(err)
		}
	}
}
