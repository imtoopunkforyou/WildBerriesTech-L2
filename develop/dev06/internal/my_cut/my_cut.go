package cut

import (
	"bufio"
	fls "dev06/internal/flags"
	"os"
	"strconv"
	"strings"
)

// NewCut создание нового cut'а.
func NewCut(fl *fls.Flags) []string {
	var result []string
	fields := getFields(fl)               // Получаем поля.
	scanner := bufio.NewScanner(os.Stdin) // Создаем сканер, считывающий stdin.
	for scanner.Scan() {
		if cutStr := cutString(scanner.Text(), fields, fl); cutStr != nil { // Вырезаем строку.
			result = append(result, *cutStr) // Пушим вырезанную строку в конечный слайс.
		}
	}

	return result
}

func cutString(str string, fields []int, fl *fls.Flags) *string {
	if *fl.S && !strings.Contains(str, *fl.D) {
		return nil // Если у нас флаг -s и нет разделителя, то выходим.
	}

	var line []string
	if strings.Contains(str, *fl.D) {
		line = strings.Split(str, *fl.D) // Разбиваем строку stdin по разделителю на слайс.
	}

	cut := cutFieldsFromLine(fields, line)

	if len(cut) == 0 {
		cut = append(cut, str)
	}

	if len(cut) > 0 {
		result := strings.Join(cut, *fl.D) // Преобразуем слайс обратно в строку для возврата результата.
		return &result
	}

	return nil
}

func cutFieldsFromLine(fields []int, line []string) []string {
	var cut []string               // Сюда буду пушиться вырезанные части.
	for _, field := range fields { // Проходимся по слайсу полей.
		if field < 0 {
			field *= -1
		}
		if field <= len(line) { // Проверка, что у нас такое поле есть в строке.
			cut = append(cut, line[field-1])
		} else {
			if len(cut) < len(line) { // Если у нас нет такого поля, то добавляем пустую строку.
				cut = append(cut, "")
			}
		}
	}
	return cut
}

// Функция для получения слайса интов с полями, которые нужно вырезать.
func getFields(fl *fls.Flags) []int {
	var columns []int
	fields := strings.Split(*fl.F, ",") // Разбиваем строку с полученными полями на слайс.
	for _, field := range fields {
		if column, err := strconv.Atoi(field); err == nil {
			columns = append(columns, column)
		} // Берем через атои, тк нам могут подать не числа.
	}
	return columns
}
