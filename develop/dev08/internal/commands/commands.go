package commands

import (
	"bytes"
	"fmt"
	"github.com/shirou/gopsutil/process"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// ForkExec вызов стандартных команд из оболочки терминала.
func ForkExec(str []string, r io.Reader, w io.Writer) {
	var outputBuffer bytes.Buffer

	cmd := exec.Command(str[0], str[1:]...) // Вызываем команды из терминала.

	cmd.Stdin = r // Назначаем вводом reader
	cmd.Stderr = os.Stderr
	cmd.Stdout = &outputBuffer // Назначаем наш буффер. Изначально был сразу writer, но если вывод последней строки
	// не содержит переноса строки, то ввод не будет перенесен на новую строку.
	if err := cmd.Run(); err != nil {
		fmt.Fprintln(w, err)
	}

	res := strings.TrimRight(outputBuffer.String(), "\n") // Избавляемся от лишней пустой строки, если он есть.
	if len(res) > 0 {                                     // Проверка, чтобы не получить пустой вывод в терминал.
		fmt.Fprintln(w, res)
	}
}

// Echo вывод аргумента в STDOUT.
func Echo(str []string) string {
	var res strings.Builder
	for _, line := range str[1:] {
		line = strings.Trim(line, "\"")
		line = strings.Trim(line, "'")
		res.WriteString(line + " ")
	}
	return res.String()
}

// Pwd показать путь до текущего каталога.
func Pwd() string {
	wd, err := os.Getwd()
	if err != nil {
		return err.Error()
	}
	return wd
}

// Kill "убить" процесс, переданный в качестве аргумента.
func Kill(str []string) {
	if len(str) < 2 {
		fmt.Println("kill: not enough arguments")
		return
	}

	for _, line := range str[1:] {
		id, err := strconv.Atoi(line) // Получаем pid из строки.
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

// Ps выводит общую информацию по запущенным процессам в формате такой-то формат.
func Ps(w io.Writer) {
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

// Cd смена директории (в качестве аргумента могут быть то-то и то).
func Cd(str []string) {
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
