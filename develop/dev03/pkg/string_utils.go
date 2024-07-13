package stringutils

import "strconv"

// RemoveLastEmptyLine Удаление последней пустой строи
func RemoveLastEmptyLine(fileLines *[]string) {
	if (*fileLines)[len(*fileLines)-1] == "" {
		*fileLines = (*fileLines)[:len(*fileLines)-1]
	} // Удаляем, чтобы она не входила в сортировку, как в оригинале.
}

// ExtractNumericStrings Извлечение числовых строк из изначального слайса.
func ExtractNumericStrings(fileLines *[]string) []float64 {
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

// DeleteDuplicates Удаление повторяющихся строк.
func DeleteDuplicates(fileLines *[]string) {
	var newLines []string

	for i := 0; i < len(*fileLines); i++ { // Проходимся по заранее отсортированном слайсу.
		if i > 0 && (*fileLines)[i] == (*fileLines)[i-1] {
			continue // Если текущие элемент и предыдущие одинаковые - переходим к следующему.
		}
		newLines = append(newLines, (*fileLines)[i]) // Добавляем элементы в новый слайс без дубликатов.
	}

	*fileLines = newLines // Присваиваем новый слайс без повторений к старому.
}
