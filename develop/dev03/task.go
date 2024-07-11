package main

/*
Отсортировать строки в файле по аналогии с консольной утилитой sort (man sort — смотрим описание и основные параметры):
на входе подается файл с несортированными строками, на выходе — файл с отсортированными.

Реализовать поддержку утилитой следующих ключей:

-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительно:

Реализовать поддержку утилитой следующих ключей:

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учетом суффиксов
*/

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type flags struct {
	n, r, u *bool
	k       *int
}

func main() {
	fl := flags{}
	flagParse(&fl)

	fileLines, err := readFile(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatalln(err)
	}

	newSort(&fileLines, &fl)
	for _, str := range fileLines {
		fmt.Println(str)
	}
}

// Сортировка по флагам.
func newSort(fileLines *[]string, fl *flags) {
	removeLastEmptyLine(fileLines)
	if *fl.k > 0 {
		columnSort(fileLines, fl.k)
	}
	if *fl.n {
		numericSort(fileLines)
	}
	if *fl.u {
		sort.Strings(*fileLines)
		deleteDuplicates(fileLines)
	}
	if *fl.r {
		reverseSort(fileLines)
	}
	if *fl.k == 0 && !*fl.n && !*fl.u && !*fl.r {
		sort.Strings(*fileLines)
	}
}

// Парсинг флагов.
func flagParse(fl *flags) {
	fl.k = flag.Int("k", 0, "Column for sort.")
	fl.n = flag.Bool("n", false, "Sort by numeric value.")
	fl.r = flag.Bool("r", false, "Reverse sort.")
	fl.u = flag.Bool("u", false, "Without duplicate string.")

	flag.Parse()
}

// Функция считывания файла.
func readFile(inputFile string) ([]string, error) {
	file, err := os.ReadFile(inputFile) // Считывание файла.
	if err != nil {
		return nil, err
	}

	str := string(file)               // Преобразуем файл в строку.
	lines := strings.Split(str, "\n") // Разбиваем преобразованный файл на строки.

	return lines, nil
}

// Сортировка строк.
func removeLastEmptyLine(fileLines *[]string) {
	if (*fileLines)[len(*fileLines)-1] == "" {
		*fileLines = (*fileLines)[:len(*fileLines)-1]
	} // Удаляем последнюю строку, если она пустая, чтобы она не входила в сортировку, как в оригинале.
}

func columnSort(fileLines *[]string, column *int) {
	*column-- // Уменьшаем полученную колонку, тк индекс идет с 0.
	sort.SliceStable(*fileLines, func(i, j int) bool {
		// Разбиение строк на подстроки-колонки.
		colsI := strings.Split((*fileLines)[i], " ")
		colsJ := strings.Split((*fileLines)[j], " ")

		// Проверяем существует ли такие колонки.
		if *column < len(colsI) && *column < len(colsJ) {
			// Если колонки одинаковые, то сортируем лексикографически.
			if colsI[*column] == colsJ[*column] {
				return (*fileLines)[i] < (*fileLines)[j]
			} // Иначе сравниваем по значениям в этих колонках.
			return colsI[*column] < colsJ[*column]
		}
		// Если колонка есть в строке J, но нет в строке I, то строка I считается меньшей и идет раньше.
		if *column >= len(colsI) && *column < len(colsJ) {
			return true
		}
		// Если колонка есть в строке I, но нет в строке J, то строка J считается меньшей и идет раньше.
		if *column < len(colsI) && *column >= len(colsJ) {
			return false
		}
		// Если ни одно условие не подошло, то сравниваем лексикографически.
		return (*fileLines)[i] < (*fileLines)[j]
	})
}

// Сортирует по числовому значению.
func numericSort(fileLines *[]string) {
	numSlice := extractNumericStrings(fileLines)
	sort.Strings(*fileLines) // Отдельно сортируем строки.
	sort.Float64s(numSlice)  // Отдельно сортируем числа.

	for _, num := range numSlice {
		*fileLines = append(*fileLines, strconv.FormatFloat(num, 'f', -1, 64))
	} // Добавляем отсортированные числа в наш изначальный слайс.
}

// Извлечение числовых строк из изначального слайса.
func extractNumericStrings(fileLines *[]string) []float64 {
	var charSlice []string // Для записи нечисловых строк.
	var numSlice []float64 // Для записи числовых строк

	for _, r := range *fileLines {
		if num, err := strconv.ParseFloat(r, 64); err == nil {
			numSlice = append(numSlice, num)
		} else {
			charSlice = append(charSlice, r)
		}
	}
	*fileLines = charSlice // Записываем нечисловые строки в изначальный слайс.

	return numSlice // Получаем слайс с числами из числовых строк.
}

// Сортировка в обратном порядке.
func reverseSort(fileLines *[]string) {
	sort.SliceStable(*fileLines, func(i, j int) bool {
		return (*fileLines)[i] > (*fileLines)[j]
		// Возвращает true если элементом под индексом i должен идти перед элементом под индексом j.
	})
}

// Удаление повторяющихся строк.
func deleteDuplicates(fileLines *[]string) {
	var newLines []string

	for i := 0; i < len(*fileLines); i++ { // Проходимся по заранее отсортированном слайсу.
		if i > 0 && (*fileLines)[i] == (*fileLines)[i-1] {
			continue // Если текущие элемент и предыдущие одинаковые - переходим к следующему.
		}
		newLines = append(newLines, (*fileLines)[i]) // Добавляем элементы в новый слайс без дубликатов.
	}

	*fileLines = newLines // Присваиваем новый слайс без повторений к старому.
}
