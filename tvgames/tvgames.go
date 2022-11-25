package tvgames

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode"

	"git.hjkl.gq/team13/birdazzone-api/model"
	"git.hjkl.gq/team13/birdazzone-api/tvgames/birdazzone"
	"git.hjkl.gq/team13/birdazzone-api/tvgames/gametracker"
	"git.hjkl.gq/team13/birdazzone-api/tvgames/ghigliottina"
	"git.hjkl.gq/team13/birdazzone-api/twitter"
	"git.hjkl.gq/team13/birdazzone-api/util"
	"github.com/gin-gonic/gin"
	geojson "github.com/paulmach/go.geojson"
	"github.com/swaggo/swag/example/celler/httputil"
)

var gameTrackers = []gametracker.GameTracker{
	ghigliottina.GetGhigliottinaTracker(),
	birdazzone.GetBirdazzoneTracker(),
}

var gameTrackersById = map[int]*gametracker.GameTracker{}

var games []model.Game

var gamesById = map[int]*model.Game{}

func TvGamesGroup(group *gin.RouterGroup) {
	group.GET("/", getTvGames)
	group.GET("/:id", getTvGameById)
	group.GET("/:id/solution", gameSolution)
	group.GET("/:id/attempts", gameAttempts)
	group.GET("/:id/attempts/stats", gameAttemptsStats)
	group.GET("/:id/results", gameResults)
}

// getTvGames godoc
// @Summary Get all TV games
// @Tags    tvgames
// @Produce json
// @Success 200 {array} model.Game
// @Router  /tvgames [get]
func getTvGames(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, games)
}

// getTvGameById godoc
// @Summary Get TV game
// @Tags    tvgames
// @Produce json
// @Param   id  path     int true "ID to search"
// @Success 200 {object} model.Game
// @Failure 404 {object} model.Error "game id not found"
// @Router  /tvgames/{id} [get]
func getTvGameById(ctx *gin.Context) {
	game, err := util.IdToObject(ctx, gamesById)
	if err == nil {
		ctx.JSON(http.StatusOK, game)
	}
}

// gameSolution godoc
// @Summary Retrieve game's solution
// @Tags    tvgames
// @Produce json
// @Param   id   path     string true  "Game to query"
// @Param   date query    string false "Date to query; if not specified, last game instance is considered" Format(date)
// @Success 200  {object} model.GameKey
// @Success 204  {string} string      "No game instance has been played"
// @Failure 400  {object} model.Error "integer parsing error (id) or error while parsing to date"
// @Failure 404  {object} model.Error "game id not found"
// @Failure 500  {object} model.Error "(internal server error)"
// @Router  /tvgames/{id}/solution [get]
func gameSolution(ctx *gin.Context) {

	gameTracker, err := util.IdToObject(ctx, gameTrackersById)
	if err != nil {
		return
	}

	date_str, hasDate := ctx.GetQuery("date")
	var date time.Time
	var sol model.GameKey
	if hasDate {
		date, err = util.StringToDate(date_str)
		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest,
				fmt.Errorf("integer parsing error (id) or error while parsing to date"))
			return
		}
		if gameTracker.Solution != nil {
			sol, err = gameTracker.Solution(date)
		} else {
			httputil.NewError(ctx, http.StatusInternalServerError,
				fmt.Errorf("missing solution getter for %T", gameTracker))
		}
	} else {
		if gameTracker.Solution != nil {
			sol, err = gameTracker.LastSolution()
		} else {
			httputil.NewError(ctx, http.StatusInternalServerError,
				fmt.Errorf("missing last solution getter for %T", gameTracker))
		}
	}

	if err == nil {
		ctx.JSON(http.StatusOK, sol)
	} else {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
	}

}

func getAttempts(ctx *gin.Context, successesOnly bool, fromStr string, toStr string) (*twitter.ProfileTweets, error) {
	fmt.Println(fromStr + " | " + toStr)
	gameTracker, err := util.IdToObject(ctx, gameTrackersById)
	if err != nil {
		return nil, err
	}
	query := gameTracker.Query

	if fromStr == "" {
		//!from && !to
		lastSol, err := gameTracker.LastSolution()
		if successesOnly {
			if err != nil {
				return nil, err
			}
			query += " " + lastSol.Key
		}
		return twitter.GetManyRecentTweetsFromQuery(query, "", lastSol.Date)
	} else {
		//from enabled
		t, _ := util.StringToDateTime(fromStr)
		sol, err := gameTracker.Solution(t)
		if successesOnly {
			if err != nil {
				return nil, err
			}
			query += " " + sol.Key
		}
		if sol.Date < toStr || toStr == "" { //from && !to
			toStr = sol.Date
		}
		return twitter.GetManyRecentTweetsFromQuery(query, fromStr, toStr)
	}
}

func toLowerAlphaOnly(r rune) rune {
	if unicode.IsLetter(r) {
		return unicode.ToLower(r)
	}
	if unicode.IsDigit(r) || strings.ContainsRune("#:/@", r) {
		return r
	}
	return ' '
}

var attemptsBlacklist = []string{
	"indovinato",
	"indovinata",
	"perché",
	"perchè",
	"perche'",
	"soluzione",
	"oppure",
	"sicuramente",
	"però",
	"peró",
	"non",
	"stasera",
	"voi",
}

func tweetTextToAttempt(text string) string {
	for _, word := range strings.Split(strings.Map(toLowerAlphaOnly, text), " ") {
		if len(word) >= 3 &&
			word[0:1] != "#" &&
			(len(word) == 3 || word[0:4] != "http") &&
			util.IsAlphabetic(&word) &&
			!util.Contains(&attemptsBlacklist, word) {
			return word
		}
	}
	return ""
}

// gameAttempts godoc
// @Summary Retrieve game's attempts
// @Tags    tvgames
// @Produce json
// @Param   id         path     string true  "Game to query"
// @Param   from       query    string false "Starting instant of the time interval used to filter the tweets. If not specified, the beginning of the last game instance is used"                                                                                    Format(date-time)
// @Param   to         query    string false "Ending instant of the time interval used to filter the tweets. Must be later than but in the same day of the starting instant. If not specified, the ending of the game happening during the starting instant is used" Format(date-time)
// @Param   pageIndex  query    int    false "Index of the page to query"                                                                                                                                                                                            minimum(1) default(1)
// @Param   pageLength query    int    false "Length of the page to query"                                                                                                                                                                                           minimum(1) default(10)
// @Success 200        {object} model.Page[model.Tweet]
// @Success 204        {string} string      "no game instance has been played"
// @Failure 400        {object} model.Error "integer parsing error (pageIndex) or pageIndex < 1 or integer parsing error (pageLength) or pageIndex < pageLength or integer parsing error (id) or error while parsing to date"
// @Failure 404        {object} model.Error "game id not found"
// @Failure 500        {object} model.Error "(internal server error)"
// @Router  /tvgames/{id}/attempts [get]
func gameAttempts(ctx *gin.Context) {
	var pageIndex, pageLength int
	pageIndex, err := strconv.Atoi(ctx.DefaultQuery("pageIndex", "1"))
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, errors.New("integer parsing error (pageIndex) or pageIndex < 1 or integer parsing error (pageLength) or pageIndex < pageLength or integer parsing error (id) or error while parsing to date"))
		return
	}
	if pageIndex < 1 {
		httputil.NewError(ctx, http.StatusBadRequest, errors.New("integer parsing error (pageIndex) or pageIndex < 1 or integer parsing error (pageLength) or pageIndex < pageLength or integer parsing error (id) or error while parsing to date"))
		return
	}
	pageLength, err = strconv.Atoi(ctx.DefaultQuery("pageLength", "1"))
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, errors.New("integer parsing error (pageIndex) or pageIndex < 1 or integer parsing error (pageLength) or pageIndex < pageLength or integer parsing error (id) or error while parsing to date"))
		return
	}
	if pageLength < 1 {
		httputil.NewError(ctx, http.StatusBadRequest, errors.New("integer parsing error (pageIndex) or pageIndex < 1 or integer parsing error (pageLength) or pageIndex < pageLength or integer parsing error (id) or error while parsing to date"))
		return
	}

	fromStr, hasFrom := ctx.GetQuery("from")
	toStr, hasTo := ctx.GetQuery("to")
	var fromTime, toTime time.Time
	if hasFrom {
		fromTime, err = util.StringToDateTime(fromStr)
		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, errors.New("integer parsing error (pageIndex) or pageIndex < 1 or integer parsing error (pageLength) or pageIndex < pageLength or integer parsing error (id) or error while parsing to date"))
			return
		}
		fromStr = util.DateToString(fromTime)
		if hasTo {
			toTime, err = util.StringToDateTime(toStr)
			if err != nil {
				httputil.NewError(ctx, http.StatusBadRequest, errors.New("integer parsing error (pageIndex) or pageIndex < 1 or integer parsing error (pageLength) or pageIndex < pageLength or integer parsing error (id) or error while parsing to date"))
				return
			}
			toStr = util.DateToString(toTime)
			if fromStr > toStr {
				httputil.NewError(ctx, http.StatusBadRequest, errors.New("integer parsing error (pageIndex) or pageIndex < 1 or integer parsing error (pageLength) or pageIndex < pageLength or integer parsing error (id) or error while parsing to date"))
				return
			}
			if toStr > util.DateToString(time.Now()) {
				httputil.NewError(ctx, http.StatusBadRequest, errors.New("integer parsing error (pageIndex) or pageIndex < 1 or integer parsing error (pageLength) or pageIndex < pageLength or integer parsing error (id) or error while parsing to date"))
				return
			}
			if fromTime.Day() != toTime.Day() || fromTime.Month() != toTime.Month() || fromTime.Year() != toTime.Year() {
				httputil.NewError(ctx, http.StatusBadRequest, errors.New("integer parsing error (pageIndex) or pageIndex < 1 or integer parsing error (pageLength) or pageIndex < pageLength or integer parsing error (id) or error while parsing to date"))
				return
			}
		}
	} else {
		fromStr = ""
		toStr = ""
	}

	result, err := getAttempts(ctx, true, fromStr, toStr)
	if err != nil {
		httputil.NewError(ctx, http.StatusNoContent, errors.New("no game instance has been played"))
		return
	}
	tweets := result.Data
	util.Reverse(&tweets)
	// users
	usersById := make(map[string]twitter.Profile, len(result.Includes.Users))
	for _, user := range result.Includes.Users {
		usersById[user.ID] = user
	}
	// places
	placesById := make(map[string]twitter.Place, len(result.Includes.Places))
	for _, place := range result.Includes.Places {
		placesById[place.ID] = place
	}
	n := len(tweets)
	from := util.Max(0, util.Min(pageLength*(pageIndex-1), n-1))
	res := make([]model.Tweet, util.Min(from+pageLength, n)-from)
	for i := range res {
		tweet := tweets[from+i]
		var geometry *geojson.Geometry
		if tweet.Geo != nil {
			geo := placesById[tweet.Geo.PlaceId].Geo
			geometry = &geo
		}
		res[i] = model.MakeTweet(tweet, usersById[tweet.AuthorId], geometry)
	}
	ctx.JSON(http.StatusOK, model.Page[model.Tweet]{Entries: res, NumberOfPages: (n + pageLength - 1) / pageLength})
}

func getAttemptsStats(ctx *gin.Context) (model.Chart, error) {
	gameTracker, err := util.IdToObject(ctx, gameTrackersById)
	if err != nil {
		return nil, err
	}
	fromStr, hasFrom := ctx.GetQuery("from")
	toStr, hasTo := ctx.GetQuery("to")
	var fromTime, toTime time.Time
	if hasFrom {
		fromTime, err = util.StringToDateTime(fromStr)
		if err != nil {
			return nil, errors.New("integer parsing error (id) or error while parsing to date")
		}
		fromStr = util.DateToString(fromTime)
		if hasTo {
			toTime, err = util.StringToDateTime(toStr)
			if err != nil {
				return nil, errors.New("integer parsing error (id) or error while parsing to date")
			}
			toStr = util.DateToString(toTime)
			if fromStr > toStr {
				return nil, errors.New("integer parsing error (id) or error while parsing to date")
			}
			if toStr > util.DateToString(time.Now()) {
				return nil, errors.New("integer parsing error (id) or error while parsing to date")
			}
			if fromTime.Day() != toTime.Day() || fromTime.Month() != toTime.Month() || fromTime.Year() != toTime.Year() {
				return nil, errors.New("integer parsing error (id) or error while parsing to date")
			}
		}
	} else {
		fromStr = ""
		toStr = ""
	}

	result, err := getAttempts(ctx, false, fromStr, toStr)
	if err != nil {
		return nil, err
	}
	tweets := result.Data
	var solution model.GameKey
	if fromStr != "" {
		s, _ := util.StringToDateTime(fromStr)
		solution, err = gameTracker.Solution(s)
	} else {
		solution, err = gameTracker.LastSolution()
	}

	if err != nil {
		return nil, err
	}
	chartAsMap := make(map[string]int)
	for _, tweet := range tweets {
		text := strings.ToLower(tweet.Text)
		var attempt string
		if strings.Contains(text, solution.Key) {
			attempt = solution.Key
		} else {
			attempt = tweetTextToAttempt(text)
		}
		if attempt == "" {
			continue
		}
		_, ok := chartAsMap[attempt]
		if ok {
			chartAsMap[attempt]++
		} else {
			chartAsMap[attempt] = 1
		}
	}
	chart := make(model.Chart, len(chartAsMap))
	i := 0
	for k, v := range chartAsMap {
		chart[i] = model.ChartEntry{Value: k, AbsoluteFrequency: v}
		i++
	}
	return chart, nil
}

// gameAttemptsStats godoc
// @Summary Retrieve game's attempts' frequencies
// @Tags    tvgames
// @Produce json
// @Param   id   path     string true  "Game to query"
// @Param   from query    string false "Starting instant of the time interval used to filter the tweets. If not specified, the beginning of the last game instance is used"                                                                                    Format(date-time)
// @Param   to   query    string false "Ending instant of the time interval used to filter the tweets. Must be later than but in the same day of the starting instant. If not specified, the ending of the game happening during the starting instant is used" Format(date-time)
// @Success 200  {object} model.Chart
// @Success 204  {string} string      "No game instance has been played"
// @Failure 400  {object} model.Error "integer parsing error (id) or error while parsing to date"
// @Failure 404  {object} model.Error "game id not found"
// @Router  /tvgames/{id}/attempts/stats [get]
func gameAttemptsStats(ctx *gin.Context) {
	chart, err := getAttemptsStats(ctx)
	if err == nil {
		ctx.JSON(http.StatusOK, chart)
	} else {
		httputil.NewError(ctx, http.StatusBadRequest, err)
	}
}

// gameResults godoc
// @Summary Retrieve game's number of successes and failures, grouped in time interval bins
// @Tags    tvgames
// @Produce json
// @Param   id   path     string             true  "Game to query"
// @Param   from query    string             false "Starting date of the time interval used to filter the tweets. If not specified, the last game instance's date is used"                             Format(date)
// @Param   to   query    string             false "Ending date of the time interval used to filter the tweets. Cannot be earlier than the starting date. If not specified, the starting date is used" Format(date)
// @Param   each query    int                false "Number of seconds for the duration of each time interval bin the retrieved tweets are to be grouped by"                                            minimum(1)
// @Success 200  {array}  model.BooleanChart "A array of boolean charts comparing successes and failures in the game. Each boolean chart is labeled as the starting instant of its time interval bin"
// @Success 204  {string} string             "No game instance has been played"
// @Failure 400  {object} model.Error        "integer parsing error (id) or date parsing error (from) or date parsing error (to) or to > today or from > to or integer parsing error (each) or each < 1"
// @Failure 404  {object} model.Error        "game id not found"
// @Router  /tvgames/{id}/results [get]
func gameResults(ctx *gin.Context) {
	// TODO: implement filter based on time
	gameTracker, err := util.IdToObject(ctx, gameTrackersById)
	if err == nil {
		var result *twitter.ProfileTweets
		result, err = getAttempts(ctx, false, "", "")
		if err == nil {
			tweets := result.Data
			var solution model.GameKey
			solution, err = gameTracker.LastSolution()
			if err == nil {
				successes := 0
				for _, tweet := range tweets {
					if strings.Contains(strings.ToLower(tweet.Text), solution.Key) {
						successes++
					}
				}
				ctx.JSON(
					http.StatusOK,
					model.BooleanChart{
						Positives: successes,
						Negatives: len(tweets) - successes,
					},
				)
				return
			}
		}
	}
	httputil.NewError(ctx, http.StatusInternalServerError, err)
}

func assignIds() {
	for i := range gameTrackers {
		gameTrackers[i].Game.Id = i
	}
}

func initDataStructures() {
	games = make([]model.Game, len(gameTrackers))
	i := 0
	for k := range gameTrackers {
		v := &gameTrackers[k]
		gameTrackersById[k] = v
		games[i] = v.Game
		gamesById[k] = &(v.Game)
		i += 1
	}
}

func init() {
	assignIds()
	initDataStructures()
}
