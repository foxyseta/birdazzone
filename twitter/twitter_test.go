package twitter

import (
	"net/http"
	"testing"
	"time"

	"git.hjkl.gq/team13/birdazzone-api/util"
)

func TestGetTweetsFromQuery(t *testing.T) {
	res := getTweetsFromQuery("kalulu")
	if res.Code != http.StatusOK {
		t.Fatalf("Expected code %d but instead got %d", http.StatusOK, res.Code)
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

func TestGetUser(t *testing.T) {
	uid := getUser("quizzettone")
	if uid == nil {
		t.Fatal("Error in GetUser")
	}
}

func TestGetTweetsFromUser(t *testing.T) {
	tweets := getTweetsFromUser("1499992669480755204", "")
	if tweets == nil {
		t.Fatal("Error in GetTweetsFromUser")
	}
}

func TestGetGhigliottinaSolution(t *testing.T) {
	sol := getGhigliottinaSolution()
	if sol == "ERR_USER" {
		t.Fatal("Error in retrieving solution's author")
	}
	if sol == "ERR_TWEETS" {
		t.Fatal("Error in retrieving solution's tweets")
	}
	if sol == "" {
		t.Log("No available solution")
	}
}

func TestTimeFormat(t *testing.T) {
	tm := time.Date(2022, time.September, 10, 18, 59, 00, 00, time.UTC)
	x := timeFormat(tm)
	if x != "2022-09-09T18:55:00Z" {
		t.Fatal("Error in formatting time #1")
	}
	tm = time.Date(2022, time.December, 01, 10, 19, 01, 00, time.UTC)
	x = timeFormat(tm)
	if x != "2022-11-30T18:55:00Z" {
		t.Fatal("Error in formatting time #2")
	}
}
