package tvgames

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetTvGames(t *testing.T) {
	//should return all games + Code 200
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	getTvGames(c)
	fmt.Printf("%d %s\n", w.Code, w.Body.String())
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}
func TestGetTvGameById(t *testing.T) {

	//id:0 should return something + Code 200
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "0"}}
	getTvGameById(c)
	fmt.Printf("%d %s\n", w.Code, w.Body.String())
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

	//predicted failure (id:-1 should return Code 400)
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "-1"}}
	getTvGameById(c)
	fmt.Printf("%d %s\n", w.Code, w.Body.String())
	if w.Code != http.StatusBadRequest {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusBadRequest, w.Code)
	}

}
func TestInitAPI(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	v1 := r.Group("/api/v1")
	InitAPI(v1)

	//test on /tvgames
	req, err := http.NewRequest(http.MethodGet, "/api/v1/tvgames", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

	//test on /tvgames/:id
	req, err = http.NewRequest(http.MethodGet, "/api/v1/tvgames/0", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

}
