package main

import (
	"bufio"
	"dev08/internal/app"
	"dev08/internal/commands"
	"fmt"
	"os"
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
		wd := strings.Split(commands.Pwd(), "/") // Для вывода конкретной директории разбиваем весь путь на элементы слайса.
		fmt.Print(">/", wd[len(wd)-1], " ")      // Выводим конкретную директорию, в которой мы находимся

		reader := bufio.NewScanner(os.Stdin) // Создаем reader из stdin.
		if reader.Scan() {                   // Сканируем полученные данные из stdin.
			shell.HandleLinuxPipes(reader.Text())
		}
	}
}
