package main_test

import (
	fls "dev06/internal/flags"
	c "dev06/internal/my_cut"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func execSourceCut(input string, isFile bool, flags []string) ([]string, error) {
	var cut *exec.Cmd

	cut = exec.Command("cut", flags...)
	if isFile {
		file, err := os.Open(input)
		if err != nil {
			return nil, err
		}

		cut.Stdin = file
	} else {
		cut.Stdin = strings.NewReader(input)
	}

	output, err := cut.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("cut: %v\n%s", err, string(output))
	}

	lines := strings.Split(string(output), "\n")

	return lines[:len(lines)-1], nil // Убираем таким образом пустую строку из stdin.
}

func execMyCut(input string, isFile bool, fl *fls.Flags) ([]string, error) {
	originalStdin := os.Stdin
	if isFile {
		file, err := os.Open(input)
		if err != nil {
			return nil, err
		}

		os.Stdin = file
	} else {
		r, w, _ := os.Pipe()

		if _, err := w.Write([]byte(input)); err != nil {
			return nil, err
		}

		if err := w.Close(); err != nil {
			return nil, err
		}

		os.Stdin = r
	}

	res := c.NewCut(fl)
	os.Stdin = originalStdin

	return res, nil
}

func initFlags(fValue, rValue string, sValue bool) *fls.Flags {
	return &fls.Flags{F: &fValue, D: &rValue, S: &sValue}
}

func compareOutput(myCut, srcCut []string) error {
	if len(myCut) != len(srcCut) {
		return fmt.Errorf("count strings myLine: %d, srcLine: %d", len(myCut), len(srcCut))
	}
	for i := 0; i < len(myCut); i++ {
		if len(myCut[i]) != len(srcCut[i]) && myCut[i] != srcCut[i] {
			return fmt.Errorf("myLine: %d %s, srcLine: %d %s", i+1, myCut[i], i+1, srcCut[i])
		}
	}
	return nil
}

func TestCase1(t *testing.T) {
	input := "naruto\tabob\njojo\t"

	srcCut, err := execSourceCut(input, false, []string{"-f", "1,2", "-s"})
	if err != nil {
		t.Errorf("cut failed: %v", err)
	}

	fl := initFlags("1,2", "\t", true)

	myCut, err := execMyCut(input, false, fl)
	if err != nil {
		t.Errorf("exec myCut failed: %v", err)
	}

	if err = compareOutput(myCut, srcCut); err != nil {
		t.Errorf("output failed: %v", err)
	}
}

func TestCase2(t *testing.T) {
	srcCut, err := execSourceCut("test.txt", true, []string{"-d", " ", "-f", "1,2"})
	if err != nil {
		t.Errorf("cut failed: %v", err)
	}

	fl := initFlags("1,2", " ", false)

	myCut, err := execMyCut("test.txt", true, fl)
	if err != nil {
		t.Errorf("exec myCut failed: %v", err)
	}

	if err = compareOutput(myCut, srcCut); err != nil {
		t.Errorf("output failed: %v", err)
	}
}

func TestCase3(t *testing.T) {
	srcCut, err := execSourceCut("test.txt", true, []string{"-d", ",", "-f", "1,2,3", "-s"})
	if err != nil {
		t.Errorf("cut failed: %v", err)
	}

	fl := initFlags("1,2,3", ",", true)

	myCut, err := execMyCut("test.txt", true, fl)
	if err != nil {
		t.Errorf("exec myCut failed: %v", err)
	}

	if err = compareOutput(myCut, srcCut); err != nil {
		t.Errorf("output failed: %v", err)
	}
}

func TestCase4(t *testing.T) {
	srcCut, err := execSourceCut("test.txt", true, []string{"-d", "|", "-f", "1,2,-3,4"})
	if err != nil {
		t.Errorf("cut failed: %v", err)
	}

	fl := initFlags("1,2,-3,4", "|", false)

	myCut, err := execMyCut("test.txt", true, fl)
	if err != nil {
		t.Errorf("exec myCut failed: %v", err)
	}

	//for _, line := range myCut {
	//	fmt.Println(line)
	//}
	//for _, line := range srcCut {
	//	fmt.Println(line)
	//}
	if err = compareOutput(myCut, srcCut); err != nil {
		t.Errorf("output failed: %v", err)
	}
}

func TestCase5(t *testing.T) {
	srcCut, err := execSourceCut("test.txt", true, []string{"-d", ":", "-f", "4,1", "-s"})
	if err != nil {
		t.Errorf("cut failed: %v", err)
	}

	fl := initFlags("4,1", ":", true)

	myCut, err := execMyCut("test.txt", true, fl)
	if err != nil {
		t.Errorf("exec myCut failed: %v", err)
	}

	if err = compareOutput(myCut, srcCut); err != nil {
		t.Errorf("output failed: %v", err)
	}
}

func TestCase6(t *testing.T) {
	srcCut, err := execSourceCut("test.txt", true, []string{"-f", "9"})
	if err != nil {
		t.Errorf("cut failed: %v", err)
	}

	fl := initFlags("9", "\t", false)

	myCut, err := execMyCut("test.txt", true, fl)
	if err != nil {
		t.Errorf("exec myCut failed: %v", err)
	}

	if err = compareOutput(myCut, srcCut); err != nil {
		t.Errorf("output failed: %v", err)
	}
}
