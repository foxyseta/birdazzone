package main

import(
  "net"
  "strconv"
  "strings"
  "testing"
)

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
