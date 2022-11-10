package tvgames

import (
	"net/http"
	"net/http/httptest"
	"testing"

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

func TestGetTvGameById(t *testing.T) {
	// id: 0 should return something + Code 200
	util.GetTestingGinContext().Params = []gin.Param{{Key: "id", Value: "0"}}
	getTvGameById(util.GetTestingGinContext())
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

func TestInitAPI(t *testing.T) {
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
