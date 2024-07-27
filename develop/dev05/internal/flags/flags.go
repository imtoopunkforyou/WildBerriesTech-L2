package flags

import "flag"

// Flags Структура, которая содержит в себе флаги, для различных сортировок.
type Flags struct {
	After, Before, Context                    *int
	Count, IgnoreCase, Invert, Fixed, LineNum *bool
	FirstCall                                 *bool // Нужен для разделения выводов у флагов -А -В -С
}

// FlagParse Парсинг флагов.
func FlagParse() *Flags {
	fl := Flags{}

	fl.After = flag.Int("A", 0, "Print +N lines after match")
	fl.Before = flag.Int("B", 0, "Print +N lines before match")
	fl.Context = flag.Int("C", 0, "Print ±N lines around match")
	fl.Count = flag.Bool("c", false, "Count number of lines")
	fl.IgnoreCase = flag.Bool("i", false, "Ignore case")
	fl.Invert = flag.Bool("v", false, "Invert match")
	fl.Fixed = flag.Bool("F", false, "Print not matched lines")
	fl.LineNum = flag.Bool("n", false, "Print line number")

	fl.FirstCall = new(bool)

	flag.Parse()

	return &fl
}
