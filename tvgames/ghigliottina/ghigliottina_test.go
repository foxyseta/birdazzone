package ghigliottina

import "testing"

func TestSolution(t *testing.T) {
	sol := Solution()
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
