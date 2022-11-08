package twitter

import (
	"net/http"
	"testing"

	"git.hjkl.gq/team13/birdazzone-api/util"
)

func TestGetTweetsQuery(t *testing.T) {
	res := getTweetsQuery("kalulu")
	if res.Code != 200 {
		t.Fatalf("Cannot retrieve GET request from Twitter")
	}
}

func TestReturnTweetQuery(t *testing.T) {
	body := `{"text":"TEST STRING","user":"@myusername"}`
	//should return a json + Code 200
	returnTweetQuery(util.GetTestingGinContext(), queryResponse{&body, 200, "200 OK"})
	if util.GetTestingResponseRecorder().Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d", http.StatusOK,
			util.GetTestingResponseRecorder().Code)
	}
}
