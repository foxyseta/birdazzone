package twitter

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestReturnTweetQuery(t *testing.T) {
	//should return a json + Code 200
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	returnTweetQuery(c, `{"text":"TEST STRING","user":"@myusername"}`)
	fmt.Printf("%d %s\n", w.Code, w.Body.String())
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

func TestGetRequest(t *testing.T) {
	res := getRequest("kalulu")
	if res == nil {
		t.Fatalf("Cannot retrieve GET request from Twitter")
	}

}
