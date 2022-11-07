package util

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestInit(t *testing.T) {
  if gin.Mode() != gin.TestMode {
    t.Fatal("Gin is not in test mode")
  }
  if GetTestingResponseRecorder() != testingResponseRecorder {
    t.Fatal("Wrong GetTestingResponseRecorder")
  }
  if testingGinContext == nil {
    t.Fatal("testingGinContext is nil")
  }
  if GetTestingGinContext() != testingGinContext {
    t.Fatalf("Wrong GetTestingGinContext")
  }
  if testingGinEngine == nil {
    t.Fatal("testingGinEngine is nil")
  }
  if GetTestingGinEngine() != testingGinEngine {
    t.Fatalf("Wrong GetTestingGinEngine")
  }
}
