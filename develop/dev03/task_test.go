package main

import (
	fls "dev03/internal/flags"
	s "dev03/internal/sort"
	f "dev03/pkg/file"
	"fmt"
	"os/exec"
	"strings"
	"testing"
	"unicode"
)

func execSourceSort(flags, file string) ([]string, error) {
	var sort *exec.Cmd

	if flags != "" {
		if unicode.IsDigit(rune(flags[0])) {
			sort = exec.Command("sort", "-k", flags, file)
		} else {
			sort = exec.Command("sort", flags, file)
		}
	} else {
		sort = exec.Command("sort", file)
	}

	output, err := sort.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")

	return lines, nil
}

func execMySort(fl fls.Flags, file string) ([]string, error) {
	inputFile, err := f.ReadFile(file)
	if err != nil {
		return nil, err
	}

	s.NewSort(&inputFile, &fl)

	return inputFile, nil
}

func initFlags(nValue, rValue, uValue bool, kValue int) fls.Flags {
	return fls.Flags{N: &nValue, R: &rValue, U: &uValue, K: &kValue}
}

func compareOutput(mySort, srcSort []string) error {
	if len(mySort) != len(srcSort)-1 { // -1 так как strings.Split(string(output), "\n") создает лишнюю пустую строку.
		return fmt.Errorf("count strings myLine: %d, srcLine: %d", len(mySort), len(srcSort))
	}
	for i := 0; i < len(mySort); i++ {
		if mySort[i] != srcSort[i] {
			return fmt.Errorf("myLine: %d %s, srcLine: %d %s", i+1, mySort[i], i+1, srcSort[i])
		}
	}
	return nil
}

func TestWithoutKeys(t *testing.T) {
	srcSort, err := execSourceSort("", "Makefile")
	if err != nil {
		t.Errorf("sort failed: %v", err)
	}
	fl := initFlags(false, false, false, 0)

	mySort, err := execMySort(fl, "Makefile")
	if err != nil {
		t.Errorf("readFile failed: %v", err)
	}
	err = compareOutput(mySort, srcSort)
	if err != nil {
		t.Errorf("compare failed: %v", err)
	}
}

func TestKeyU(t *testing.T) {
	srcSort, err := execSourceSort("-u", "Makefile")
	if err != nil {
		t.Errorf("sort failed: %v", err)
	}
	fl := initFlags(false, false, true, 0)

	mySort, err := execMySort(fl, "Makefile")
	if err != nil {
		t.Errorf("readFile failed: %v", err)
	}
	err = compareOutput(mySort, srcSort)
	if err != nil {
		t.Errorf("compare failed: %v", err)
	}
}

func TestKeyR(t *testing.T) {
	srcSort, err := execSourceSort("-r", "Makefile")
	if err != nil {
		t.Errorf("sort failed: %v", err)
	}
	fl := initFlags(false, true, false, 0)

	mySort, err := execMySort(fl, "Makefile")
	if err != nil {
		t.Errorf("readFile failed: %v", err)
	}
	err = compareOutput(mySort, srcSort)
	if err != nil {
		t.Errorf("compare failed: %v", err)
	}
}

func TestKeyN(t *testing.T) {
	srcSort, err := execSourceSort("-n", "Makefile")
	if err != nil {
		t.Errorf("sort failed: %v", err)
	}
	fl := initFlags(true, false, false, 0)

	mySort, err := execMySort(fl, "Makefile")
	if err != nil {
		t.Errorf("readFile failed: %v", err)
	}
	err = compareOutput(mySort, srcSort)
	if err != nil {
		t.Errorf("compare failed: %v", err)
	}
}

func TestWithCorrectColumnKeyK(t *testing.T) {
	srcSort, err := execSourceSort("1", "Makefile")
	if err != nil {
		t.Errorf("sort failed: %v", err)
	}
	fl := initFlags(false, false, false, 1)

	mySort, err := execMySort(fl, "Makefile")
	if err != nil {
		t.Errorf("readFile failed: %v", err)
	}
	err = compareOutput(mySort, srcSort)
	if err != nil {
		t.Errorf("compare failed: %v", err)
	}
}

func TestWithIncorrectColumnKeyK(t *testing.T) {
	srcSort, err := execSourceSort("2", "Makefile")
	if err != nil {
		t.Errorf("sort failed: %v", err)
	}
	fl := initFlags(false, false, false, 2)

	mySort, err := execMySort(fl, "Makefile")
	if err != nil {
		t.Errorf("readFile failed: %v", err)
	}
	err = compareOutput(mySort, srcSort)
	if err != nil {
		t.Errorf("compare failed: %v", err)
	}
}

func TestWithInvalidColumnKeyK(t *testing.T) {
	srcSort, err := execSourceSort("15", "Makefile")
	if err != nil {
		t.Errorf("sort failed: %v", err)
	}
	fl := initFlags(false, false, false, 15)

	mySort, err := execMySort(fl, "Makefile")
	if err != nil {
		t.Errorf("readFile failed: %v", err)
	}
	err = compareOutput(mySort, srcSort)
	if err != nil {
		t.Errorf("compare failed: %v", err)
	}
}
