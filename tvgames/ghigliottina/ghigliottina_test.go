package ghigliottina

import "testing"

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
	sol, err := solution()
	if err != nil {
		t.Fatal(err.Error())
	}
	if sol == "" {
		t.Fatal("Empty solution")
	}
}
