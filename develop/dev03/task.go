package main

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	fls "dev03/internal/flags"
	s "dev03/internal/sort"
	f "dev03/pkg/file"
	"fmt"
	"log"
	"os"
)

func main() {
	fileLines, err := f.ReadFile(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatalln(err)
	}

	fl := fls.FlagParse()
	s.NewSort(&fileLines, fl)

	for _, str := range fileLines {
		fmt.Println(str)
	}
}
