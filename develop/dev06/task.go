package main

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
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
