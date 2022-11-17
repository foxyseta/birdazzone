package server

import (
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
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

func TestCorsMiddleware(t *testing.T) {
	testingGinContext, _ := gin.CreateTestContext(httptest.NewRecorder())
	corsMiddleware(testingGinContext)
	header := testingGinContext.Writer.Header()
	if header.Get("Access-Control-Allow-Origin") != "*" {
		t.Fail()
	}
	if header.Get("Access-Control-Allow-Credentials") != "true" {
		t.Fail()
	}
	if header.Get("Access-Control-Allow-Headers") != "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With" {
		t.Fail()
	}
	if header.Get("Access-Control-Allow-Methods") != "POST, OPTIONS, GET, PUT" {
		t.Fail()
	}
	var err error
	testingGinContext.Request, err = http.NewRequest("OPTIONS", "https://api.twitter.com/", nil)
	if err != nil {
		t.Fatal("Unable to create request")
	}
	corsMiddleware(testingGinContext)
}

func TestAddress(t *testing.T) {
	tokens := strings.Split(Address(), ":")
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
	if len(routes) == 0 {
		t.Fatal("No routes found")
	}
}
