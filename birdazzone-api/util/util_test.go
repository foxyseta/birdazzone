package util

import (
	"testing"
	"time"

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

func TestTimeFormat(t *testing.T) {
	tm := time.Date(2022, time.September, 10, 18, 59, 00, 00, time.UTC)
	x := LastGameDate(tm)
	if x != "2022-09-09T18:55:00Z" {
		t.Fatal("Error in formatting time #1")
	}
	tm = time.Date(2022, time.December, 01, 10, 19, 01, 00, time.UTC)
	x = LastGameDate(tm)
	if x != "2022-11-30T18:55:00Z" {
		t.Fatal("Error in formatting time #2")
	}
}
