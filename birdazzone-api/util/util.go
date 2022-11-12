package util

import (
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

func (p *Pair[T, U]) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("(%T, %T)", p.First, p.Second)
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

func LastInstantAtGivenTime(dt time.Time, hours int) string {
	// before <hours> returns yesterday's solution
	if dt.Hour() < hours {
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
	x += strconv.Itoa(dt.Day()) + fmt.Sprintf("T%d:00:00Z", hours)
	return x
}

func IdToObject[T any](ctx *gin.Context, data map[int]T) (T, error) {
	key, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatus(400)
		var result T
		return result, err
	}

	value, ok := data[key]

	if !ok {
		ctx.AbortWithStatus(404)
		var result T
		return result, fmt.Errorf("Object not found using key %d", key)
	}

	return value, err
}

func QueryParamToPositiveInt(ctx *gin.Context, paramName string, defaultValue string) (int, error) {
	value, err := strconv.Atoi(ctx.DefaultQuery(paramName, defaultValue))
	if err != nil {
		newError := fmt.Errorf("integer parsing error (%s)", paramName)
		httputil.NewError(ctx, http.StatusBadRequest, newError)
		return 0, newError
	}
	if value < 1 {
		newError := fmt.Errorf("%s < 1", paramName)
		httputil.NewError(ctx, http.StatusBadRequest, newError)
		return 0, newError
	}
	return value, nil
}

func init() {
	gin.SetMode(gin.TestMode)
	testingGinContext, testingGinEngine = gin.CreateTestContext(testingResponseRecorder)
}
