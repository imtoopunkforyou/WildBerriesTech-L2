package grep

import (
	fls "dev05/internal/flags"
	"fmt"
	"log"
	"regexp"
	"strings"
)

// NewGrep Для создания нового грепа.
func NewGrep(fl fls.Flags, pattern string, fileLines []string) {
	var countLines int
	for i, line := range fileLines {
		match := findString(pattern, line, &fl)
		if *fl.Invert {
			match = !match
		}
		if match {
			if *fl.LineNum {
				fmt.Printf("%d:", i+1)
			}
			if !*fl.Count {
				printMatches(fileLines, i, &fl)
			}
			countLines++
		}
	}
	if *fl.Count {
		fmt.Println(countLines)
	}
}

func printMatches(lines []string, index int, fl *fls.Flags) {
	if (*fl.After > 0 || *fl.Before > 0) && *fl.FirstCall {
		fmt.Println("--")
	}
	*fl.FirstCall = true

	if *fl.Context > 0 {
		*fl.Before = *fl.Context
		*fl.After = *fl.Context
	}

	if *fl.Before > 0 {
		printBeforeAfterMatch(index-*fl.Before, index-1, lines)
	}

	fmt.Println(lines[index])

	if *fl.After > 0 {
		printBeforeAfterMatch(index+1, index+*fl.After, lines)
	}
}

func printBeforeAfterMatch(start, end int, lines []string) {
	for i := start; i <= end; i++ {
		if i > 0 && i < len(lines) {
			fmt.Println(lines[i])
		}
	}
}

func findString(pattern string, line string, fl *fls.Flags) bool {
	if *fl.IgnoreCase {
		pattern = strings.ToLower(pattern)
		line = strings.ToLower(line)
	}
	if *fl.Fixed {
		return pattern == line
	}
	match, err := regexp.MatchString(pattern, line)
	if err != nil {
		log.Fatal(err)
	}
	return match
}
