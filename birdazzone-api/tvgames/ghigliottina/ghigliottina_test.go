package ghigliottina

import "testing"

func TestSolution(t *testing.T) {
	sol, err := solution()
	if err != nil {
		t.Fatal(err.Error())
	}
	if sol == "" {
		t.Fatal("Empty solution")
	}
}
