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

	fileStrs, err := readFile(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatalln(err)
	}

	newSort(&fileStrs, &fl)
	for _, str := range fileStrs {
		fmt.Println(str)
	}
}

// Сортировка по флагам.
func newSort(fileStrs *[]string, fl *flags) {
	if *fl.k > 0 { // todo
		columnSort(fileStrs)
	}
	if *fl.n {
		numericSort(fileStrs)
	}
	if *fl.u {
		lexicographicalSort(fileStrs)
		deleteDuplicates(fileStrs)
	}
	if *fl.r {
		reverseSort(fileStrs)

	}
	if *fl.k == 0 && !*fl.n && !*fl.u && !*fl.r {
		lexicographicalSort(fileStrs)
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

	str := string(file)              // Преобразуем файл в строку.
	strs := strings.Split(str, "\n") // Разбиваем преобразованный файл на строки.

	return strs, nil
}

// Сортировка строк.
func lexicographicalSort(fileStrs *[]string) {
	if (*fileStrs)[len(*fileStrs)-1] == "" {
		*fileStrs = (*fileStrs)[:len(*fileStrs)-1]
	} // Удаляем последнюю строку, если она пустая, чтобы она не входила в сортировку, как в оригинале.

	sort.Strings(*fileStrs)
}

// Сортирует по колонке.
func columnSort(fileStrs *[]string) {}

// Сортирует по числовому значению.
func numericSort(fileStrs *[]string) {
	numSlice := extractNumericStrings(fileStrs)
	lexicographicalSort(fileStrs) // Отдельно сортируем строки.
	sort.Float64s(numSlice)       // Отдельно сортируем числа.

	for _, num := range numSlice {
		*fileStrs = append(*fileStrs, strconv.FormatFloat(num, 'f', -1, 64))
	} // Добавляем отсортированные числа в наш изначальный слайс.
}

// Извлечение числовых строк из изначального слайса.
func extractNumericStrings(fileStrs *[]string) []float64 {
	var charSlice []string // Для записи нечисловых строк.
	var numSlice []float64 // Для записи числовых строк

	for _, r := range *fileStrs {
		if num, err := strconv.ParseFloat(r, 64); err == nil {
			numSlice = append(numSlice, num)
		} else {
			charSlice = append(charSlice, r)
		}
	}
	*fileStrs = charSlice // Записываем нечисловые строки в изначальный слайс.

	return numSlice // Получаем слайс с числами из числовых строк.
}

// Сортировка в обратном порядке.
func reverseSort(fileStrs *[]string) {
	sort.SliceStable(*fileStrs, func(i, j int) bool {
		return (*fileStrs)[i] > (*fileStrs)[j]
		// Возвращает true если элементом под индексом i должен идти перед элементом под индексом j.
	})
}

// Удаление повторяющихся строк.
func deleteDuplicates(fileStrs *[]string) {
	var newStrs []string

	for i := 0; i < len(*fileStrs); i++ { // Проходимся по заранее отсортированном слайсу.
		if i > 0 && (*fileStrs)[i] == (*fileStrs)[i-1] {
			continue // Если текущие элемент и предыдущие одинаковые - переходим к следующему.
		}
		newStrs = append(newStrs, (*fileStrs)[i]) // Добавляем элементы в новый слайс без дубликатов.
	}

	*fileStrs = newStrs // Присваиваем новый слайс без повторений к старому.
}
