package flags

import "flag"

// Flags Структура, которая содержит в себе флаги, для различных сортировок.
type Flags struct {
	N, R, U *bool
	K       *int
}

// FlagParse Парсинг флагов.
func FlagParse() *Flags {
	fl := Flags{}
	fl.K = flag.Int("k", 0, "Column for sort.")
	fl.N = flag.Bool("n", false, "Sort by numeric value.")
	fl.R = flag.Bool("r", false, "Reverse sort.")
	fl.U = flag.Bool("u", false, "Without duplicate string.")

	flag.Parse()
	return &fl
}
