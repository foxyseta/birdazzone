package tvgames

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"unicode"

	"git.hjkl.gq/team13/birdazzone-api/util"
	"github.com/gin-gonic/gin"
)

func TestGetTvGames(t *testing.T) {
	// should return all games + Code 200
	getTvGames(util.GetTestingGinContext())
	if util.GetTestingResponseRecorder().Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, util.GetTestingResponseRecorder().Code)
	}
}

func testAPICall(t *testing.T, call func(*gin.Context)) {
	// id: 0 should return something + Code 200
	util.GetTestingGinContext().Params = []gin.Param{{Key: "id", Value: "0"}}
	call(util.GetTestingGinContext())
	if util.GetTestingResponseRecorder().Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, util.GetTestingResponseRecorder().Code)
	}
	// predicted failure (id: -1 should return Code 400)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "-1"}}
	getTvGameById(c)
	if w.Code != http.StatusNotFound {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusNotFound, w.Code)
	}
}

func TestGetTvGameById(t *testing.T) {
	testAPICall(t, getTvGameById)
}

func TestGameSolution(t *testing.T) {
	testAPICall(t, gameSolution)
}

func testGetAttempts(t *testing.T, successesOnly bool) {
	util.GetTestingGinContext().Params = []gin.Param{{Key: "id", Value: "0"}}
	result, err := getAttempts(util.GetTestingGinContext(), successesOnly)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil || result.Data == nil {
		t.Fatal("Nil result")
	}
}

func TestGetAttemptsWrongParam(t *testing.T) {
	util.GetTestingGinContext().Params = []gin.Param{{Key: "id", Value: "zebra"}}
	_, err := getAttempts(util.GetTestingGinContext(), true)
	if err == nil {
		t.FailNow()
	}
}

func TestGetAttemptsSuccessesOnly(t *testing.T) {
	testGetAttempts(t, true)
}

func TestGetAttemptsFailuresToo(t *testing.T) {
	testGetAttempts(t, false)
}

func TestToLowerAlphaOnlyWithLowercaseLetters(t *testing.T) {
	for r := 'a'; r <= 'z'; r++ {
		if toLowerAlphaOnly(r) != r {
			t.Fatalf("%c => %c", r, toLowerAlphaOnly(r))
		}
	}
}

func TestToLowerAlphaOnlyWithUppercaseLetters(t *testing.T) {
	for r := 'A'; r <= 'Z'; r++ {
		if toLowerAlphaOnly(r) != unicode.ToLower(r) {
			t.Fatalf("%c => %c", r, toLowerAlphaOnly(r))
		}
	}
}

func TestToLowerAlphaOnlyWithDigits(t *testing.T) {
	for r := '0'; r <= '9'; r++ {
		if toLowerAlphaOnly(r) != r {
			t.Fatalf("%c => %c", r, toLowerAlphaOnly(r))
		}
	}
}

func TestToLowerAlphaOnlyWithSocialSymbols(t *testing.T) {
	for _, r := range "#:/@" {
		if toLowerAlphaOnly(r) != r {
			t.Fatalf("%c => %c", r, toLowerAlphaOnly(r))
		}
	}
}

func TestToLowerAlphaOnlyWithIrrelevantSymbols(t *testing.T) {
	for _, r := range "!?&-" {
		if toLowerAlphaOnly(r) != ' ' {
			t.Fatalf("%c => %c", r, toLowerAlphaOnly(r))
		}
	}
}

func TestTweetTextToAttemptEmpty(t *testing.T) {
	if tweetTextToAttempt("") != "" {
		t.FailNow()
	}
}

func TestTweetTextToAttemptA(t *testing.T) {
	if tweetTextToAttempt("#leredità è ponte? ...") != "ponte" {
		t.FailNow()
	}
}

func TestTweetTextToAttemptB(t *testing.T) {
	if tweetTextToAttempt("è pollo mi sa!! #ovvio #ghigliottina #leredità") != "pollo" {
		t.FailNow()
	}
}

func TestGameAttempts(t *testing.T) {
	testAPICall(t, gameAttempts)
}

func TestGameAttemptsStats(t *testing.T) {
	testAPICall(t, gameAttemptsStats)
}

func TestGameResults(t *testing.T) {
	testAPICall(t, gameResults)
}

func TestInit(t *testing.T) {
	v1 := util.GetTestingGinEngine().Group("/api/v1")
	TvGamesGroup(v1)

	//test on /tvgames
	req, err := http.NewRequest(http.MethodGet, "/api/v1/tvgames", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	util.GetTestingGinEngine().ServeHTTP(util.GetTestingResponseRecorder(), req)
	if util.GetTestingResponseRecorder().Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK,
			util.GetTestingResponseRecorder().Code)
	}

	//test on /tvgames/:id
	req, err = http.NewRequest(http.MethodGet, "/api/v1/tvgames/0", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	util.GetTestingGinEngine().ServeHTTP(util.GetTestingResponseRecorder(), req)
	if util.GetTestingResponseRecorder().Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK,
			util.GetTestingResponseRecorder().Code)
	}
}
