package chess

import (
	"net/http"
	"testing"

	"git.hjkl.gq/team13/birdazzone-api/twitter"
)

func TestUncheckedGetCheckMoveWrongDate(t *testing.T) {
	_, err := uncheckedGetCheckMove("anything", "2017-09-12T20:00:00Z", 2)
	if err == nil {
		t.Fatal("Error expected")
	}
	if err.Code != http.StatusNotFound {
		t.Fatalf("Got %d, but expected 404", err.Code)
	}
}

func TestUncheckedGetCheckMoveNoQuotes(t *testing.T) {
	_, err := uncheckedGetCheckMove("twitter", "2017-09-12T20:00:00Z", 2)
	if err == nil {
		t.Fatal("Error expected")
	}
	if err.Code != http.StatusNotFound {
		t.Fatalf("Got %d, but expected 404", err.Code)
	}
}

func TestMostPopularChessMove(t *testing.T) {
	answer := mostPopularChessMove(twitter.ProfileTweet{ID: "1605757925951242240"})
	if answer != "e7e6" {
		t.Fatalf("Expected 'e7e6'. Found '%s' instead", answer)
	}
}
