package util

import (
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestReverse(t *testing.T) {
	s := []int{4, 3, 2, 1, 0}
	Reverse(&s)
	for e, i := range s {
		if e != i {
			t.Fatalf("%d found but %d expected in position #%d", e, i, i)
		}
	}
}

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
	x := LastInstantAtGivenTime(tm, 14)
	if x != "2022-09-09T10:00:00Z" {
		t.Fatal("Error in formatting time #1")
	}
	tm = time.Date(2022, time.December, 01, 10, 19, 01, 00, time.UTC)
	x = LastInstantAtGivenTime(tm, 20)
	if x != "2022-11-30T20:00:00Z" {
		t.Fatal("Error in formatting time #2")
	}
}
