package main

import (
	"os"
	"syscall"
	"testing"
)

func TestMain(m *testing.M) {
	sttyArgs := syscall.ProcAttr{
		Dir:   "",
		Env:   []string{},
		Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()},
		Sys:   nil,
	}
	syscall.ForkExec("go", []string{"run", "."}, &sttyArgs)
	os.Exit(m.Run())
}
