package util

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
)

var testingResponseRecorder = httptest.NewRecorder()
var testingGinContext *gin.Context = nil
var testingGinEngine *gin.Engine = nil

const NilRepresentation = "<nil>"

type Pair[T any, U any] struct {
	First  T
	Second U
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Contains[T comparable](array *[]T, value T) bool {
	for _, x := range *array {
		if x == value {
			return true
		}
	}
	return false
}

func IsAlphabetic(str *string) bool {
	for _, r := range *str {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func Reverse[T any](array *[]T) {
	n := len(*array)
	halfN := n / 2
	for i, j := 0, n-1; i < halfN; i++ {
		x := &(*array)[i]
		y := &(*array)[j]
		t := *x
		*x = *y
		*y = t
		j--
	}
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

func StringToDate(d string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", d)
	return t, err
}

func DateToString(d time.Time) string {
	return d.UTC().Format(time.RFC3339)
}

func LastInstantAtGivenTime(dt time.Time, hours int) string {
	// before <hours> returns yesterday's solution
	if dt.Hour() < hours {
		dt = dt.AddDate(0, 0, -1)
	}
	dt = time.Date(dt.Year(), dt.Month(), dt.Day(), hours, 0, 0, 0, &time.Location{})
	return dt.UTC().Format(time.RFC3339)
}

func IdToObject[T any](ctx *gin.Context, data map[int]T) (T, error) {
	key, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, errors.New("integer parsing error (id)"))
		var result T
		return result, err
	}

	value, ok := data[key]

	if !ok {
		httputil.NewError(ctx, http.StatusNotFound, errors.New("game id not found"))
		var result T
		return result, fmt.Errorf("object not found using key %d", key)
	}

	return value, err
}

func init() {
	gin.SetMode(gin.TestMode)
	testingGinContext, testingGinEngine = gin.CreateTestContext(testingResponseRecorder)
}
