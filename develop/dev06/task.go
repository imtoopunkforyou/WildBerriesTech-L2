package main

/*
Реализовать утилиту аналог консольной команды cut (man cut). Утилита должна принимать строки через STDIN,
разбивать по разделителю (TAB) на колонки и выводить запрошенные.

Реализовать поддержку утилитой следующих ключей:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
*/

import (
	fls "dev06/internal/flags"
	c "dev06/internal/my_cut"
	"fmt"
	"log"
)

func main() {
	fl := fls.FlagParse()
	if *fl.D == "\n" || *fl.D == "" || *fl.D == "::" { // Валидация разделителя, как в оригинале.
		log.Fatal("cut: bad delimiter")
	}

	result := c.NewCut(fl)       // Получаем вырезанные строки.
	for _, res := range result { // Выводим результат.
		fmt.Println(res)
	}
}
