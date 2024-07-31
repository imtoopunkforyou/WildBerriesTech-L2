package file

import (
	"os"
	"strings"
)

// ReadFile Функция считывания файла.
func ReadFile(inputFile string) ([]string, error) {
	file, err := os.ReadFile(inputFile) // Считывание файла.
	if err != nil {
		return nil, err
	}

	str := string(file)               // Преобразуем файл в строку.
	lines := strings.Split(str, "\n") // Разбиваем преобразованный файл на строки.

	return lines, nil
}
