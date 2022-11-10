package util

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
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
