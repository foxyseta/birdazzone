package util

import (
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

var testingResponseRecorder = httptest.NewRecorder()
var testingGinContext *gin.Context = nil

func init() {
  gin.SetMode(gin.TestMode)
  testingGinContext, _ = gin.CreateTestContext(testingResponseRecorder)
}

func GetTestingResponseRecorder() *httptest.ResponseRecorder {
  return testingResponseRecorder
}

func GetTestingGinContext() *gin.Context {
  return testingGinContext
}
