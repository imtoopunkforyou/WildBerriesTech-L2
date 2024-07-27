package main

import (
	fls "dev05/internal/flags"
	grep "dev05/internal/my_grep"
	f "dev05/pkg"
	"log"
	"os"
)

/*
Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).

Реализовать поддержку утилитой следующих ключей:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", напечатать номер строки
*/

func main() {
	fileLines, err := f.ReadFile(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatal(err)
	}
	pattern := os.Args[len(os.Args)-2]
	fl := fls.FlagParse()

	grep.NewGrep(*fl, pattern, fileLines)
}
