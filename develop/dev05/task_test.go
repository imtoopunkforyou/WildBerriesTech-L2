package main_test

import (
	"bytes"
	fls "dev05/internal/flags"
	g "dev05/internal/my_grep"
	f "dev05/pkg/file"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"
	"unicode"
)

func execSourceGrep(flags, pattern, file string) ([]string, error) {
	var grep *exec.Cmd

	if flags != "" {
		if len(flags) == 4 && unicode.IsDigit(rune(flags[3])) {
			grep = exec.Command("grep", flags, pattern, file)
		} else {
			grep = exec.Command("grep", flags, pattern, file)
		}
	} else {
		grep = exec.Command("grep", pattern, file)
	}

	output, err := grep.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")

	return lines, nil
}

func execMyGrep(fl fls.Flags, pattern, file string) ([]string, error) {
	inputFile, err := f.ReadFile(file)
	if err != nil {
		return nil, err
	}

	r, w, _ := os.Pipe()
	originalStdout := os.Stdout
	os.Stdout = w

	g.NewGrep(&fl, &pattern, &inputFile)
	if err = w.Close(); err != nil {
		return nil, err
	}
	os.Stdout = originalStdout

	var buff bytes.Buffer
	if _, err = io.Copy(&buff, r); err != nil {
		return nil, err
	}
	lines := strings.Split(string(buff.Bytes()), "\n")
	return lines, nil
}

func initFlags(AValue, BValue, CValue int, cValue, iValue, vValue, FValue, nValue bool) fls.Flags {
	return fls.Flags{After: &AValue, Before: &BValue, Context: &CValue, Count: &cValue, IgnoreCase: &iValue,
		Invert: &vValue, Fixed: &FValue, LineNum: &nValue, FirstCall: new(bool)}
}

func compareOutput(myGrep, srcGrep []string) error {
	if len(myGrep) != len(srcGrep) {
		return fmt.Errorf("count strings myLine: %d, srcLine: %d", len(myGrep), len(srcGrep))
	}
	for i := 0; i < len(myGrep); i++ {
		if len(myGrep[i]) != len(srcGrep[i]) && myGrep[i] != srcGrep[i] {
			return fmt.Errorf("myLine: %d %s, srcLine: %d %s", i+1, myGrep[i], i+1, srcGrep[i])
		}
	}
	return nil
}

func TestKeyA(t *testing.T) {
	srcGrep, err := execSourceGrep("-A 3", "main", "task.go")
	if err != nil {
		t.Errorf("grep failed: %v", err)
	}

	fl := initFlags(3, 0, 0, false, false, false, false, false)
	myGrep, err := execMyGrep(fl, "main", "task.go")
	if err != nil {
		t.Errorf("execMyGrep failed: %v", err)
	}

	err = compareOutput(myGrep, srcGrep)
	if err != nil {
		t.Errorf("compareOutput failed: %v", err)
	}
}

func TestKeyB(t *testing.T) {
	srcGrep, err := execSourceGrep("-B 3", "main", "task.go")
	if err != nil {
		t.Errorf("grep failed: %v", err)
	}

	fl := initFlags(0, 3, 0, false, false, false, false, false)
	myGrep, err := execMyGrep(fl, "main", "task.go")
	if err != nil {
		t.Errorf("execMyGrep failed: %v", err)
	}

	err = compareOutput(myGrep, srcGrep)
	if err != nil {
		t.Errorf("compareOutput failed: %v", err)
	}
}

func TestKeyContext(t *testing.T) {
	srcGrep, err := execSourceGrep("-C 3", "main", "task.go")
	if err != nil {
		t.Errorf("grep failed: %v", err)
	}

	fl := initFlags(0, 0, 3, false, false, false, false, false)
	myGrep, err := execMyGrep(fl, "main", "task.go")
	if err != nil {
		t.Errorf("execMyGrep failed: %v", err)
	}

	err = compareOutput(myGrep, srcGrep)
	if err != nil {
		t.Errorf("compareOutput failed: %v", err)
	}
}

func TestKeyCount(t *testing.T) {
	srcGrep, err := execSourceGrep("-c", "main", "task.go")
	if err != nil {
		t.Errorf("grep failed: %v", err)
	}

	fl := initFlags(0, 0, 0, true, false, false, false, false)
	myGrep, err := execMyGrep(fl, "main", "task.go")
	if err != nil {
		t.Errorf("execMyGrep failed: %v", err)
	}

	err = compareOutput(myGrep, srcGrep)
	if err != nil {
		t.Errorf("compareOutput failed: %v", err)
	}
}

func TestKeyI(t *testing.T) {
	srcGrep, err := execSourceGrep("-i", "main", "task.go")
	if err != nil {
		t.Errorf("grep failed: %v", err)
	}

	fl := initFlags(0, 0, 0, false, true, false, false, false)
	myGrep, err := execMyGrep(fl, "main", "task.go")
	if err != nil {
		t.Errorf("execMyGrep failed: %v", err)
	}

	err = compareOutput(myGrep, srcGrep)
	if err != nil {
		t.Errorf("compareOutput failed: %v", err)
	}
}

func TestKeyV(t *testing.T) {
	srcGrep, err := execSourceGrep("-v", "main", "task.go")
	if err != nil {
		t.Errorf("grep failed: %v", err)
	}

	fl := initFlags(0, 0, 0, false, false, true, false, false)
	myGrep, err := execMyGrep(fl, "main", "task.go")
	if err != nil {
		t.Errorf("execMyGrep failed: %v", err)
	}

	err = compareOutput(myGrep, srcGrep)
	if err != nil {
		t.Errorf("compareOutput failed: %v", err)
	}
}

func TestKeyF(t *testing.T) {
	srcGrep, err := execSourceGrep("-F", "package main", "task.go")
	if err != nil {
		t.Errorf("grep failed: %v", err)
	}

	fl := initFlags(0, 0, 0, false, false, false, true, false)
	myGrep, err := execMyGrep(fl, "package main", "task.go")
	if err != nil {
		t.Errorf("execMyGrep failed: %v", err)
	}

	err = compareOutput(myGrep, srcGrep)
	if err != nil {
		t.Errorf("compareOutput failed: %v", err)
	}
}

func TestKeyN(t *testing.T) {
	srcGrep, err := execSourceGrep("-n", "main", "task.go")
	if err != nil {
		t.Errorf("grep failed: %v", err)
	}

	fl := initFlags(0, 0, 0, false, false, false, false, true)
	myGrep, err := execMyGrep(fl, "main", "task.go")
	if err != nil {
		t.Errorf("execMyGrep failed: %v", err)
	}

	err = compareOutput(myGrep, srcGrep)
	if err != nil {
		t.Errorf("compareOutput failed: %v", err)
	}
}
