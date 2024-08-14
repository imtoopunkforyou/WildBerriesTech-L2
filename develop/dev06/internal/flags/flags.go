package flags

import (
	"flag"
	"log"
)

// Flags Структура, которая содержит в себе флаги, для различных сортировок.
type Flags struct {
	F, D *string
	S    *bool
}

// FlagParse Парсинг флагов.
func FlagParse() *Flags {
	fl := Flags{}

	fl.F = flag.String("f", "", "choose column for cut")
	fl.D = flag.String("d", "\t", "use another delimiter")
	fl.S = flag.Bool("s", false, "only strings with delimiter")

	flag.Parse()
	if *fl.F == "" {
		log.Fatal("myCut: option requires an argument -- f")
	}

	return &fl
}
