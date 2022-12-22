package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"git.hjkl.gq/team13/birdazzone-api/util"
	"github.com/gin-gonic/gin"
)

func TestHello(t *testing.T) {
	//should test message + Code 200

	helloWorld(util.GetTestingGinContext())
	if util.GetTestingResponseRecorder().Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK,
			util.GetTestingResponseRecorder().Code)
	}
}

func TestExistingUser(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "username", Value: "birdazzone"}}
	user(c)
	if util.GetTestingResponseRecorder().Code != http.StatusOK {
		t.Fatal(util.GetTestingResponseRecorder().Result().Status)
	}
}

func TestBirdazzoneServer(t *testing.T) {
	s := birdazzoneServer()
	if s == nil {
		t.Fatal("Cannot create server")
	}
	routes := s.Routes()
	if len(routes) == 0 {
		t.Fatal("No routes")
	}
}
