package chess

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"git.hjkl.gq/team13/birdazzone-api/twitter"
	"github.com/gin-gonic/gin"
)

func TestGetChessMoveWrongUsername(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{
		{Key: "username", Value: "@#$!"},
		{Key: "game", Value: time.Now().AddDate(-1, 0, 0).Format(time.RFC3339)},
		{Key: "turn", Value: "2"},
	}
	getChessMove(c)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("Expected 400, got %d", w.Code)
	}
}

func TestGetChessMoveWrongTurn(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{
		{Key: "username", Value: "birdazzone"},
		{Key: "game", Value: time.Now().AddDate(-1, 0, 0).Format(time.RFC3339)},
		{Key: "turn", Value: "giraffe"},
	}
	getChessMove(c)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("Expected 400, got %d", w.Code)
	}
}

func TestGetChessMoveWrongTurn2(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{
		{Key: "username", Value: "birdazzone"},
		{Key: "game", Value: time.Now().AddDate(-1, 0, 0).Format(time.RFC3339)},
		{Key: "turn", Value: "-3"},
	}
	getChessMove(c)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("Expected 400, got %d", w.Code)
	}
}

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
