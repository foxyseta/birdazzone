package gametracker

import (
	"testing"

	"git.hjkl.gq/team13/birdazzone-api/model"
)

func TestStringNull(t *testing.T) {
	var gt *GameTracker = nil
	if gt.String() != "<nil>" {
		t.Fatalf("Null GameTracker pointer is printed as '%s' instead of '<nil>'",
			gt.String())
	}
}

func TestString(t *testing.T) {
	gt := GameTracker{
		Game:  model.Game{Id: 5, Name: "SampleGame", Hashtag: "samplegame"},
		Query: "#samplegame",
	}
	if gt.String() != gt.Query {
		t.Fatalf("GameTracker for #samplegame is not printed as #samplegame")
	}
}
