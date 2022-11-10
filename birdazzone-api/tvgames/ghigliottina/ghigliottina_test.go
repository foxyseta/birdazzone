package ghigliottina

import "testing"

func TestSolution(t *testing.T) {
	sol, err := Solution()
	if err != nil {
		t.Fatal(err.Error())
	}
	if sol == "" {
		t.Fatal("Empty solution")
	}
}
