package grep

import (
	fls "dev05/internal/flags"
	"fmt"
	"log"
	"regexp"
	"strings"
)

// NewGrep Для создания грепа.
func NewGrep(fl *fls.Flags, pattern *string, fileLines *[]string) {

	var countLines int // Счетчик найденных совпадений.
	for i, line := range *fileLines {
		match := findString(*pattern, line, fl) // Находим строку.
		if *fl.Invert {
			match = !match // Инвертируем результат поиска для флага -v.
		}
		if match {
			if *fl.LineNum { // Добавляем номер найденной строки для флага -n.
				fmt.Printf("%d:", i+1)
			}
			if !*fl.Count { // Чтобы не выводило найденные совпадения для флага -c.
				printMatches(*fileLines, i, fl)
			}
			countLines++
		}
	}
	if *fl.Count { // Выводим только количество совпадений для флага -c.
		fmt.Println(countLines)
	}
}

// Функция для вывода найденных совпадений (и соседних строк для флагов -A -B -C).
func printMatches(lines []string, index int, fl *fls.Flags) {
	if (*fl.After > 0 || *fl.Before > 0) && *fl.FirstCall {
		fmt.Println("--")
	} // Для разделения совпадений, если выводи еще соседние строки.
	*fl.FirstCall = true // Устанавливаем флаг, чтобы условие выше срабатывало после первого вызова.

	if *fl.Context > 0 { // Если флаг -С, то вызываем оба флага -А -В.
		*fl.Before = *fl.Context
		*fl.After = *fl.Context
	}

	if *fl.Before > 0 { // Вывод строк до совпадения для флагов -В -С.
		printBeforeAfterMatch(index-*fl.Before, index-1, lines)
	}

	if !(index == len(lines)-1 && lines[index] == "") { // Условие чтобы не выводило последнюю пустую строку.
		fmt.Println(lines[index])
	}

	if *fl.After > 0 { // Вывод строк до совпадения для флагов -А -С.
		printBeforeAfterMatch(index+1, index+*fl.After, lines)
	}
}

// Функция, которая выводит соседние строки от найденных совпадений (для флагов -А -В -С).
func printBeforeAfterMatch(start, end int, lines []string) {
	for i := start; i <= end; i++ {
		if i > 0 && i < len(lines) {
			fmt.Println(lines[i])
		}
	}
}

// Поиск нужной строки с помощью регулярного выражения.
func findString(pattern string, line string, fl *fls.Flags) bool {
	if *fl.IgnoreCase { // Если флаг -i, то приводим все к единому регистру.
		pattern = strings.ToLower(pattern)
		line = strings.ToLower(line)
	}
	if *fl.Fixed { // Если флаг -F, то сравниваем строку с нашим шаблоном напрямую, не доходя до регулярного выражения.
		return pattern == line
	}
	match, err := regexp.MatchString(pattern, line)
	if err != nil {
		log.Fatal(err)
	}
	return match
}
