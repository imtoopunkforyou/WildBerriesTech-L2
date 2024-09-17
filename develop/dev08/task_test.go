package main_test

import (
	"bytes"
	shell "dev08/internal/app"
	"dev08/internal/commands"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

//func execSourceCmd(commands []string) ([]string, error) {
//	var cut *exec.Cmd
//
//	cut = exec.Command(commands[0], commands[1:]...)
//
//	output, err := cut.CombinedOutput()
//	if err != nil {
//		return nil, fmt.Errorf("command: %v\n%s", err, string(output))
//	}
//
//	lines := strings.Split(string(output), "\n")
//
//	return lines[:len(lines)-1], nil // Убираем таким образом пустую строку из stdin.
//}

func TestEchoQuotes(t *testing.T) {
	expect := "aboba"

	res := commands.Echo(strings.Fields("echo \"aboba\""))
	if res != expect {
		t.Errorf("expected: %v len: %v \n, got: %v len: %v", expect, len(expect), res, len(res))
	}

	res = commands.Echo(strings.Fields("echo 'aboba'"))
	if res != expect {
		t.Errorf("expected: %v len: %v \n, got: %v len: %v", expect, len(expect), res, len(res))
	}

	res = commands.Echo(strings.Fields("echo aboba"))
	if res != expect {
		t.Errorf("expected: %v len: %v \n, got: %v len: %v", expect, len(expect), res, len(res))
	}
}

func TestEchoText(t *testing.T) {
	expect := "aboba\nlol\tsus"

	res := commands.Echo(strings.Fields("echo aboba\\nlol\\tsus"))
	if res != expect {
		t.Errorf("expected: %v len: %v,\ngot: %v len: %v", expect, len(expect), res, len(res))
	}

	emptyEcho := commands.Echo(strings.Fields("echo \"\""))
	if emptyEcho != "" {
		t.Errorf("expected: %v len: %v,\ngot: %v len: %v", expect, len(expect), res, len(res))
	}
}

func TestPWD(t *testing.T) {
	var res bytes.Buffer

	expect, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	expect += "\n" // Потому что используется Fprintln.

	shell.Execution([]string{"pwd"}, nil, &res)
	if res.String() != expect {
		t.Errorf("expected: %v len: %v,\ngot: %v len: %v", expect, len(expect), res.String(), len(res.String()))
	}
}

func TestPS(t *testing.T) {
	var res bytes.Buffer

	shell.Execution([]string{"pwd"}, nil, &res)
	if len(res.String()) == 0 {
		t.Errorf("got nothing from ps command")
	}
}

func TestForkLS(t *testing.T) {
	var res bytes.Buffer
	expect := "Makefile\ngo.mod\ngo.sum\ninternal\ntask.go\ntask_test.go\n"

	shell.Execution([]string{"ls"}, nil, &res)
	if res.String() != expect {
		t.Errorf("expected: %v len: %v,\ngot: %v len: %v", expect, len(expect), res.String(), len(res.String()))
	}
}

func TestCD(t *testing.T) {
	shell.Execution([]string{"cd"}, nil, nil)

	wd, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Error(err)
	}
	if wd != homeDir {
		t.Errorf("expected change to user home directory")
	}
}

func TestKill(t *testing.T) {
	cmd := exec.Command("sleep", "10")
	err := cmd.Start()
	if err != nil {
		t.Fatalf("err to start process: %v", err)
	}

	commands.Kill([]string{"kill", strconv.Itoa(cmd.Process.Pid)})

	if err = cmd.Wait(); err == nil {
		t.Errorf("expected process to be killed, but it is still running")
	}
}
