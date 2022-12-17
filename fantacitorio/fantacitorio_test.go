package fantacitorio

import (
	"net/http"
	"testing"

	"git.hjkl.gq/team13/birdazzone-api/util"
)

func TestGetPoliticians(t *testing.T) {
	getPoliticians(util.GetTestingGinContext())
	if util.GetTestingResponseRecorder().Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, util.GetTestingResponseRecorder().Code)
	}
}

func TestGetTeams(t *testing.T) {
	getTeams(util.GetTestingGinContext())
	if util.GetTestingResponseRecorder().Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, util.GetTestingResponseRecorder().Code)
	}
}
