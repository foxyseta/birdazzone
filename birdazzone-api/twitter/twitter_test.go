package twitter

import (
	"testing"
	"time"

	"git.hjkl.gq/team13/birdazzone-api/util"
)

func TestGetUser(t *testing.T) {
	_, err := GetUser("quizzettone")
	if err != nil {
		t.Fatal("Error in GetUser")
	}
}

func TestGetTweetsFromUser(t *testing.T) {
	_, err := GetTweetsFromUser("1499992669480755204", 10, util.LastInstantAtGivenTime(time.Now(), 22))
	if err != nil {
		t.Fatal(err)
	}
}
