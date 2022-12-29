package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
)

func getToken() string {
	token, ok := os.LookupEnv("TOKEN")
	if !ok {
		fmt.Fprintf(os.Stderr, "Error: TOKEN environment variable not found.\n")
		os.Exit(1)
	}
	return token
}

var BearerToken = getToken()

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

func StringToDateTime(d string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, d)
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

func GetRequest(
	urlTemplate string,
	twitterToken bool,
	pathParams []any,
	queryParams ...Pair[string, string],
) ([]byte, error) {
	if pathParams == nil {
		pathParams = []any{}
	}
	client := &http.Client{}

	// path parameters
	for i := range pathParams {
		pathParams[i] = url.PathEscape(fmt.Sprint(pathParams[i]))
	}
	req, err := http.NewRequest("GET", fmt.Sprintf(urlTemplate, pathParams...), nil)
	if err != nil {
		return nil, err
	}
	if twitterToken {
		req.Header.Set("Authorization", "Bearer "+BearerToken)
	}

	// query parameters
	if queryParams != nil {
		q := req.URL.Query()
		for _, queryParam := range queryParams {
			if queryParam.Second != "" {
				q.Add(
					queryParam.First,
					queryParam.Second,
				)
			}
		}
		req.URL.RawQuery = q.Encode()
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("from Twitter API: %s (%s) -> (%s)", resp.Status, req.URL.String(), string(body))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return []byte(string(body)), nil
}

func init() {
	gin.SetMode(gin.TestMode)
	testingGinContext, testingGinEngine = gin.CreateTestContext(testingResponseRecorder)
}
