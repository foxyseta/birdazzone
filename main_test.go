package main

import(
  "net"
  "os"
  "strconv"
  "strings"
  "testing"
)

func TestLookupEnvWithFallbackSuccess(t *testing.T) {
  const VarName = "FOOD"
  const VarValue = "PIZZA"
  os.Setenv(VarName,VarValue)
  if lookupEnvWithFallback(VarName, "") != os.Getenv(VarName) {
    t.Fatal("Wrong variable value")
  }
}

func TestLookupEnvWithFallbackFailure(t *testing.T) {
  const VarName = "DRINK"
  const Fallback = "WATER"
  os.Unsetenv(VarName)
  if lookupEnvWithFallback(VarName, Fallback) != Fallback {
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

func TestMain(m *testing.M) {
  runServer()
  os.Exit(m.Run())
}
