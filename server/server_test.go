package server

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"syscall"
	"testing"
)

func TestLookupEnvWithFallbackFailure(t *testing.T) {
	const VarName = "DRINK"
	const Fallback = "WATER"
	os.Unsetenv(VarName)
	if lookupEnvWithFallback(VarName, Fallback) != Fallback {
		t.Fatal("Wrong variable value")
	}
}
func TestLookupEnvWithFallbackSuccess(t *testing.T) {
	const VarName = "FOOD"
	const VarValue = "PIZZA"
	os.Setenv(VarName, VarValue)
	if lookupEnvWithFallback(VarName, "") != os.Getenv(VarName) {
		t.Fatal("Wrong variable value")
	}
}

func TestAddress(t *testing.T) {
	tokens := strings.Split(address(), ":")
	if len(tokens) != 2 {
		t.Fatal("Wrong number of colons")
	}
	if net.ParseIP(tokens[0]) == nil {
		t.Fatal(tokens[0] + " is no valid host")
	}
	i, err := strconv.Atoi(tokens[1])
	if err != nil {
		t.Fatal(tokens[1] + " is no integer number")
	}
	if i < 0 {
		t.Fatalf("Negative port number %d", i)
	}
}

func TestCreate(t *testing.T) {
	s := CreateServer()
	if s == nil {
		t.Fatal("Cannot create server")
	}
	routes := s.Routes()
	if len(routes) != 1 {
		t.Fatal("Number of initial routes should be 1")
	}
	docs := routes[0]
	if docs.Method != "GET" {
		t.Fatalf("Method is %s and should be GET", docs.Method)
	}
	if docs.Path != "/swagger/*any" {
		t.Fatalf("Path is %s and should be /swagger/*any", docs.Path)
	}
}

func TestRun(t *testing.T) {
	resp, err := http.Get("http://" + address() + "/swagger/index.html")
	if err != nil {
		t.Fatalf("HTTP Error: %s", err.Error())
	}
	if resp.StatusCode != 200 {
		t.Fatalf("Status Code: got %d instead of 200", resp.StatusCode)
	}
}

func TestMain(m *testing.M) {
	//util.TestWithServer(m)
	sttyArgs := syscall.ProcAttr{
		Dir:   "",
		Env:   []string{},
		Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()},
		Sys:   nil,
	}

	pid, err := syscall.ForkExec("/bin/go", []string{"/bin/go", "run", ""}, &sttyArgs)
	fmt.Println(pid)
	fmt.Println(err)

	os.Exit(m.Run())
}
