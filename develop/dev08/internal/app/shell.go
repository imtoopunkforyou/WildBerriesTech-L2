package shell

import (
	"dev08/internal/commands"
	"fmt"
	"io"
	"os"
	"strings"
)

// HandleLinuxPipes обрабатывает ввод, в зависимости пришла нам одна команда или коневеер.
func HandleLinuxPipes(input string) {
	if len(input) > 0 {
		cmd := strings.Split(input, "|") // Разбиваем инпут на пайпы.

		var prevReader io.Reader = nil // Начальный поток ввода.
		for i, command := range cmd {
			if strings.TrimSpace(command) != "" { // Проверяем не пустая ли команда в пайпе.
				commandSlice := strings.Fields(command)

				if i == len(cmd)-1 { // Проверяем последняя/единственная ли команда.
					Execution(commandSlice, prevReader, os.Stdout)
				} else {
					pr, pw := io.Pipe() // Создаем пайп для передачи данных между командами.
					go func(cmd []string, in io.Reader, out io.Writer) {
						Execution(cmd, in, out)
						pw.Close() // Закрываем writer после выполнения команды.
					}(commandSlice, prevReader, pw)

					prevReader = pr // Обновляем поток ввода для следующей команды
				}
			}
		}
	}
}

// Execution Вызов команд для терминала.
func Execution(str []string, r io.Reader, w io.Writer) {
	switch str[0] { // Смотрим какая команда пришла.
	case "pwd":
		fmt.Fprintln(w, commands.Pwd())
	case "echo":
		fmt.Fprintln(w, commands.Echo(str))
	case "kill":
		commands.Kill(str)
	case "ps":
		commands.Ps(w)
	case "cd":
		commands.Cd(str)
	case "\\exit":
		os.Exit(0)
	default:
		commands.ForkExec(str, r, w)
	}
}
