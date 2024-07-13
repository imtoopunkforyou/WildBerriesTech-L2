package sort

import (
	fls "dev03/internal/flags"
	utl "dev03/pkg"
	"sort"
	"strconv"
	"strings"
)

// NewSort Сортировка по флагам.
func NewSort(fileLines *[]string, fl *fls.Flags) {

	utl.RemoveLastEmptyLine(fileLines)
	if *fl.K > 0 {
		columnSort(fileLines, fl.K)
	}
	if *fl.N {
		numericSort(fileLines)
	}
	if *fl.U {
		sort.Strings(*fileLines)
		utl.DeleteDuplicates(fileLines)
	}
	if *fl.R {
		reverseSort(fileLines)
	}
	if *fl.K == 0 && !*fl.N && !*fl.U && !*fl.R {
		sort.Strings(*fileLines)
	}
}

// Сортировка по колонкам.
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
	numSlice := utl.ExtractNumericStrings(fileLines)
	sort.Strings(*fileLines) // Отдельно сортируем строки.
	sort.Float64s(numSlice)  // Отдельно сортируем числа.

	for _, num := range numSlice {
		*fileLines = append(*fileLines, strconv.FormatFloat(num, 'f', -1, 64))
	} // Добавляем отсортированные числа в наш изначальный слайс.
}

// Сортировка в обратном порядке.
func reverseSort(fileLines *[]string) {
	sort.SliceStable(*fileLines, func(i, j int) bool {
		return (*fileLines)[i] > (*fileLines)[j]
		// Возвращает true если элементом под индексом i должен идти перед элементом под индексом j.
	})
}
