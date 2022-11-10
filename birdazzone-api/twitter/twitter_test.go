package twitter

import (
	"testing"
)

func TestGetUser(t *testing.T) {
	uid := GetUser("quizzettone")
	if uid == nil {
		t.Fatal("Error in GetUser")
	}
}

func TestGetTweetsFromUser(t *testing.T) {
	tweets := GetTweetsFromUser("1499992669480755204", "")
	if tweets == nil {
		t.Fatal("Error in GetTweetsFromUser")
	}
}
