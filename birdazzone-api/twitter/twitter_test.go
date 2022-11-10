package twitter

import (
	"testing"
)

func TestGetUser(t *testing.T) {
	_, err := GetUser("quizzettone")
	if err != nil {
		t.Fatal("Error in GetUser")
	}
}

func TestGetTweetsFromUser(t *testing.T) {
	_, err := GetTweetsFromUser("1499992669480755204", 10, "")
	if err != nil {
		t.Fatal("Error in GetTweetsFromUser")
	}
}
