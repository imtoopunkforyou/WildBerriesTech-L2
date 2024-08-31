package main

import (
	fls "dev05/internal/flags"
	grep "dev05/internal/my_grep"
	f "dev05/pkg/file"
	"log"
	"os"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	fileLines, err := f.ReadFile(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatal(err)
	}
	pattern := os.Args[len(os.Args)-2]
	fl := fls.FlagParse()
	grep.NewGrep(fl, &pattern, &fileLines)
}
