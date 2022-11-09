package main

import (
	"net/http"
	"testing"

	"git.hjkl.gq/team13/birdazzone-api/util"
)

func TestHello(t *testing.T) {
	//should test message + Code 200

	helloWorld(util.GetTestingGinContext())
	if util.GetTestingResponseRecorder().Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK,
			util.GetTestingResponseRecorder().Code)
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
