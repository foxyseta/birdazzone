package util

import (
	"net/http/httptest"

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
