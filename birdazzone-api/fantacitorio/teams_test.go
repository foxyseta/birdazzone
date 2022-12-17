package fantacitorio

import (
	"net/http"
	"testing"
)

func TestTeamsFromTwitter(t *testing.T) {
	_, err := teamsFromTwitter("", true)
	if err.Code != http.StatusBadRequest {
		t.Fatalf("Expected code %d but got %d", http.StatusBadRequest, err.Code)
	}
	_, err = teamsFromTwitter("elonmusk", true) //immagino che Elon Musk non giocherà mai al fantacitorio
	if err.Code != http.StatusNotFound {
		t.Fatalf("Expected code %d but got %d", http.StatusNotFound, err.Code)
	}
	_, err = teamsFromTwitter("", false)
	if err.Message != "" {
		t.Fatalf("Expected smooth run but got error %d", err.Code)
	}
}
