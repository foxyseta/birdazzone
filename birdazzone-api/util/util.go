package util

import (
	"net/http/httptest"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var testingResponseRecorder = httptest.NewRecorder()
var testingGinContext *gin.Context = nil
var testingGinEngine *gin.Engine = nil

func init() {
	gin.SetMode(gin.TestMode)
	testingGinContext, testingGinEngine = gin.CreateTestContext(testingResponseRecorder)
}

func GetTestingResponseRecorder() *httptest.ResponseRecorder {
	return testingResponseRecorder
}

func GetTestingGinContext() *gin.Context {
	return testingGinContext
}

func GetTestingGinEngine() *gin.Engine {
	return testingGinEngine
}

func TimeFormat(dt time.Time) string {
	// before 7PM returns yesterday's solution
	if dt.Hour() < 19 {
		dt = dt.AddDate(0, 0, -1)
	}
	// time format YYYY-MM-DDTHH:MM:SSZ (ISO 8601/RFC 3339 UTC TIMEZONE)
	x := strconv.Itoa(dt.Year()) + "-"
	if int(dt.Month()) < 10 {
		x += "0"
	}
	x += strconv.Itoa(int(dt.Month())) + "-"
	if dt.Day() < 10 {
		x += "0"
	}
	x += strconv.Itoa(dt.Day()) + "T18:55:00Z"
	return x
}
