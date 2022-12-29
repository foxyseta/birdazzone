package ghigliottina

import (
	"testing"
	"time"
)

func TestGhihliottinaTracker(t *testing.T) {
	res := GetGhigliottinaTracker()
	if res.Game != ghigliottinaTracker.Game {
		t.Fatal("Wrong game")
	}
	if res.Query != ghigliottinaTracker.Query {
		t.Fatal("Wrong query")
	}
}

func TestSolution(t *testing.T) {
	// checking last solution
	sol, err := lastSolution()
	if err != nil {
		t.Fatal(err.Error())
	}
	if sol.Key == "" {
		t.Fatal("Empty solution #1")
	}
	now := time.Now()
	var tm time.Time
	// checking yesterday's solution
	// tm = now.AddDate(0, 0, -1)
	// sol, err = givenSolution(tm)
	// if err != nil {
	// 	t.Fatal(err.Error())
	// }
	// if sol.Key == "" {
	// 	t.Fatal("Empty solution #2")
	// }
	// checking 7 days ago solution
	// tm = now.AddDate(0, 0, -6)
	// sol, err = givenSolution(tm)
	// if err != nil {
	// 	t.Fatal(err.Error())
	// }
	// if sol.Key == "" {
	// 	t.Fatal("Empty solution #3")
	// }
	// checking tomorrow's solution
	tm = now.AddDate(0, 0, 1)
	_, err = givenSolution(tm)
	if err == nil {
		t.Fatal("Didn't get expected error")
	}
}
