package twitter

import (
	"net/http"
	"testing"

	"git.hjkl.gq/team13/birdazzone-api/util"
)

func TestReturnTweetQuery(t *testing.T) {
	//should return a json + Code 200
	returnTweetQuery(util.GetTestingGinContext(),
    `{"text":"TEST STRING","user":"@myusername"}`)
	if util.GetTestingResponseRecorder().Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d", http.StatusOK,
      util.GetTestingResponseRecorder().Code)
	}
}

func TestGetRequest(t *testing.T) {
	res := getRequest("kalulu")
	if res == nil {
		t.Fatalf("Cannot retrieve GET request from Twitter")
	}
}
