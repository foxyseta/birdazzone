package gametracker

import (
	"testing"
	"time"

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

func pseudoSolution(a string, b string) (model.GameKey, error) {
	if a == "" && b == "" {
		return model.GameKey{Key: "key", Date: "date"}, nil
	}
	return model.GameKey{}, nil
}

func TestGivenSolution(t *testing.T) {
	_, err := GivenSolution(time.Now(), pseudoSolution)
	if err != nil {
		t.Fatalf("Operation failure")
	}
}

func TestLastSolution(t *testing.T) {
	gk, err := LastSolution(pseudoSolution)
	if err != nil {
		t.Fatalf("Operation failure")
	}
	if gk.Key != "key" {
		t.Fatalf("Wrong key")
	}
	if gk.Date != "date" {
		t.Fatalf("Wrong date")
	}
}

func TestMakeGameKeySuccess(t *testing.T) {
	gk, err := MakeGameKey("Game", "key", "2022-12-01")
	if err != nil {
		t.Fatalf("Operation failure")
	}
	if gk.Key != "key" {
		t.Fatalf("Incorrect key '%s'", gk.Key)
	}
	if gk.Date != "2022-12-01" {
		t.Fatalf("Incorrect date '%s'", gk.Date)
	}
}

func TestMakeGameKeyFailure(t *testing.T) {
	_, err := MakeGameKey("Game", "", "2022-12-01")
	if err == nil {
		t.Fatalf("Error expected")
	}
}
