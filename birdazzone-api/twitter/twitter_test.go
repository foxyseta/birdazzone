package twitter

import (
	"net/http"
	"testing"

	"git.hjkl.gq/team13/birdazzone-api/util"
)

func TestGetTweetsFromQuery(t *testing.T) {
	res := getTweetsFromQuery("kalulu")
	if res.Code != 200 {
		t.Fatalf("Cannot retrieve GET request from Twitter")
	}
}

func TestReturnTweetQuery(t *testing.T) {
	body := `{"text":"TEST STRING","user":"@myusername"}`
	//should return a json + Code 200
	returnTweetQuery(util.GetTestingGinContext(), QueryResponse{&body, 200, "200 OK"})
	if util.GetTestingResponseRecorder().Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d", http.StatusOK,
			util.GetTestingResponseRecorder().Code)
	}
}
