package unpack

import (
	"fmt"
	"strconv"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:

"a4bc2d5e" => "aaaabccddddde"
"abcd" => "abcd"
"45" => "" (некорректная строка)
"" => ""
Дополнительно:
Реализовать поддержку escape-последовательностей. Например:

qwe\4\5 => qwe45 (*)
qwe\45 => qwe44444 (*)
qwe\\5 => qwe\\\\\ (*)
В случае если была передана некорректная строка, функция должна возвращать ошибку. Написать unit-тесты.
*/

// StringUnpack Функция, которая распаковывает строку.
func StringUnpack(str string) (string, error) {
	s := []rune(str)

	if str != "" && unicode.IsDigit(s[0]) {
		return "", fmt.Errorf("invalid string: cannot start with digit: %s", string(s[0]))
	} // Проверяем, что не пустая строка начинается с символа.

	var (
		newStr []rune
		numStr string
		char   rune
	)

	for i, r := range s {
		if unicode.IsDigit(r) {
			if numStr == "" { // Проверка, чтобы у нас символ не перезаписался на число.
				char = s[i-1] // Записываем символ, который перед числом.
			}
			numStr += string(s[i]) // Для записи чисел двухзначных и более чисел.
			if i == len(str)-1 {   // Для обработки последнего символа.
				symbolsUnpack(&newStr, &numStr, char)
			}
		} else {
			symbolsUnpack(&newStr, &numStr, char)
			newStr = append(newStr, r) // Отдельный append нужен для символов, которые не повторяются больше одного раза.
		}
	}

	return string(newStr), nil
}

// Функция, которая распаковывает символ.
func symbolsUnpack(newStr *[]rune, numStr *string, char rune) {
	if *numStr != "" {
		num, err := strconv.Atoi(*numStr) // Считывание числа из строки.
		if err != nil {
			fmt.Println("error:", err)
			newStr = nil
			return
		}

		for i := 1; i < num; i++ { // Начало цикла с 1, так как в конце функции unpack мы добавляем еще раз.
			*newStr = append(*newStr, char)
		}

		*numStr = "" // После считывания числа, мы обнуляем строку для перезаписи.
	}
}
