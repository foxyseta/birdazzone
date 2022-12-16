package tvgames

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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
	call(c)
	if w.Code != http.StatusNotFound {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusNotFound, w.Code)
	}
}

func TestGameAttemptsWParams(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{{Key: "id", Value: "0"}}
	//ti := time.Now().AddDate(0, 0, -1)
	fmt.Println(c.Request)

	//c.Request.URL, _ = url.Parse("?from=" + util.DateToString(time.Date(ti.Year(), ti.Month(), ti.Day(), 0, 0, 0, 0, time.UTC)))
	//c.Request.URL.Query().Set("from", util.DateToString(time.Date(ti.Year(), ti.Month(), ti.Day(), 0, 0, 0, 0, time.UTC)))
	gameAttempts(c)
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

func TestGetTvGameById(t *testing.T) {
	testAPICall(t, getTvGameById)
}

func TestGameSolution(t *testing.T) {
	testAPICall(t, gameSolution)
}

func testGetAttempts(t *testing.T, successesOnly bool) {
	//id=0
	util.GetTestingGinContext().Params = []gin.Param{{Key: "id", Value: "0"}}
	result, err := getAttempts(util.GetTestingGinContext(), successesOnly, "", "")
	if err != nil {
		t.Fatal(err)
	}
	if result == nil || result.Data == nil {
		t.Fatal("Nil result")
	}
	//id=-1
	util.GetTestingGinContext().Params = []gin.Param{{Key: "id", Value: "-1"}}
	_, err = getAttempts(util.GetTestingGinContext(), successesOnly, "", "")
	if err == nil {
		t.Fatal("Didn't get expected error")
	}
	//test from
	d := time.Now().AddDate(0, 0, -1)
	util.GetTestingGinContext().Params = []gin.Param{{Key: "id", Value: "0"}}
	result, err = getAttempts(util.GetTestingGinContext(), successesOnly, util.DateToString(d), "")
	if err != nil {
		if err.Error() == "couldn't find Ghigliottina solution" {
			return // may happen where no game has been played
		}
		t.Fatalf(err.Error())
	}
	if result == nil || result.Data == nil {
		t.Fatal("Nil result")
	}
	//test from and to
	util.GetTestingGinContext().Params = []gin.Param{{Key: "id", Value: "0"}}
	result, err = getAttempts(util.GetTestingGinContext(), successesOnly, util.DateToString(d), util.DateToString(time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 0, 0, time.UTC)))
	if err != nil {
		t.Fatal(err)
	}
	if result == nil || result.Data == nil {
		t.Fatal("Nil result")
	}

}

func TestGetAttemptsWrongParam(t *testing.T) {
	util.GetTestingGinContext().Params = []gin.Param{{Key: "id", Value: "zebra"}}
	_, err := getAttempts(util.GetTestingGinContext(), true, "", "")
	if err == nil {
		t.FailNow()
	}
}

func TestExtractGameResultsTimes(t *testing.T) {
	util.GetTestingGinContext().Params = []gin.Param{{Key: "id", Value: "0"}}

	gameTracker, _ := util.IdToObject(util.GetTestingGinContext(), gameTrackersById)
	// should give errors
	_, _, err := extractGameResultsTimes(gameTracker, "", "")
	if err.Code != http.StatusBadRequest {
		t.Fatalf("Expected code %d but got %d", http.StatusBadRequest, err.Code)
	}
	_, _, err = extractGameResultsTimes(gameTracker, "2021-12-15", "randomtext")
	if err.Code != http.StatusBadRequest {
		t.Fatalf("Expected code %d but got %d", http.StatusBadRequest, err.Code)
	}
	_, _, err = extractGameResultsTimes(gameTracker, "2021-12-15", "2021-12-14")
	if err.Code != http.StatusBadRequest {
		t.Fatalf("Expected code %d but got %d", http.StatusBadRequest, err.Code)
	}
	_, _, err = extractGameResultsTimes(gameTracker, "2021-12-15", "2099-12-15")
	if err.Code != http.StatusBadRequest {
		t.Fatalf("Expected code %d but got %d", http.StatusBadRequest, err.Code)
	}
	// should be ok
	_, _, err = extractGameResultsTimes(gameTracker, "2021-12-15", "")
	if err.Message != "" {
		t.Fatalf("Expected OK but got %d", err.Code)
	}
	_, _, err = extractGameResultsTimes(gameTracker, "2021-12-15", "2021-12-18")
	if err.Message != "" {
		t.Fatalf("Expected OK but got %d", err.Code)
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
