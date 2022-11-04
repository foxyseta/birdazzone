package util

import (
	"fmt"
	"os"
	"syscall"
	"testing"
)

// TODO:
func TestWithServer(m *testing.M) {
	sttyArgs := syscall.ProcAttr{
		Dir:   "",
		Env:   []string{},
		Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()},
		Sys:   nil,
	}

	pid, err := syscall.ForkExec("/bin/go", []string{"/bin/go", "run", ""}, &sttyArgs)
	fmt.Println(pid)
	fmt.Println(err)
	for i := 0; true; i-- {

	}

	os.Exit(m.Run())
}
