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
	"git.hjkl.gq/team13/birdazzone-api/tvgames/gametracker"
	"git.hjkl.gq/team13/birdazzone-api/tvgames/ghigliottina"
	"git.hjkl.gq/team13/birdazzone-api/twitter"
	"git.hjkl.gq/team13/birdazzone-api/util"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
)

var gameTrackers = []gametracker.GameTracker{
	ghigliottina.GetGhigliottinaTracker(),
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
// @Summary     Get all TV games
// @Tags        tvgames
// @Produce     json
// @Success     200 {array} model.Game
// @Router      /tvgames	[get]
func getTvGames(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, games)
}

// getTvGameById godoc
// @Summary     Get TV game
// @Tags        tvgames
// @Produce     json
// @Param       id	path	int	true	"ID to search"
// @Success     200 {object} model.Game
// @Failure     404	{object}	model.Error "game id not found"
// @Router      /tvgames/{id} [get]
func getTvGameById(ctx *gin.Context) {
	game, err := util.IdToObject(ctx, gamesById)
	if err == nil {
		ctx.JSON(http.StatusOK, game)
	}
}

// gameSolution godoc
// @Summary     Retrieve game's solution
// @Tags        tvgames
// @Produce     json
// @Param       id	path	string	true	"Game to query"
// @Success     200 {string} string
// @Failure     400	{object}	model.Error "integer parsing error (id)"
// @Failure     404	{object}	model.Error "game id not found"
// @Failure     500	{object}	model.Error "(internal server error)"
// @Router      /tvgames/{id}/solution [get]
func gameSolution(ctx *gin.Context) {
	gameTracker, err := util.IdToObject(ctx, gameTrackersById)
	if err != nil {
		return
	}
	solution := gameTracker.Solution
	if solution != nil {
		s, err := gameTracker.Solution()
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{"solution": s})
		} else {
			httputil.NewError(ctx, http.StatusInternalServerError, err)
		}
	} else {
		httputil.NewError(ctx, http.StatusInternalServerError,
			fmt.Errorf("missing solution getter for %T", gameTracker))
	}
}

func getAttempts(ctx *gin.Context, successesOnly bool) (*twitter.ProfileTweets, error) {
	gameTracker, err := util.IdToObject(ctx, gameTrackersById)
	if err != nil {
		return nil, err
	}
	query := gameTracker.Query
	if successesOnly {
		solution, err := gameTracker.Solution()
		if err != nil {
			return nil, err
		}
		query += " " + solution
	}
	return twitter.GetManyRecentTweetsFromQuery(query, util.LastInstantAtGivenTime(time.Now(), 18))
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
// @Summary     Retrieve game's attempts
// @Tags        tvgames
// @Produce     json
// @Param       id path string true "Game to query"
// @Param       pageIndex query int false "Number of the page to query" minimum(1) default(1)
// @Param       pageLength query int false "Length of the page to query" minimum(1) default(10)
// @Success     200 {object} model.Page[model.Tweet]
// @Failure     400	{object}	model.Error "integer parsing error (pageIndex)"
// @Failure     400	{object}	model.Error "pageIndex < 1"
// @Failure     400	{object}	model.Error "integer parsing error (pageLength)"
// @Failure     400	{object}	model.Error "pageIndex < pageLength"
// @Failure     400	{object}	model.Error "integer parsing error (id)"
// @Failure     404	{object}	model.Error "game id not found"
// @Failure     500	{object}	model.Error "(internal server error)"
// @Router      /tvgames/{id}/attempts [get]
func gameAttempts(ctx *gin.Context) {
	var pageIndex, pageLength int
	pageIndex, err := strconv.Atoi(ctx.DefaultQuery("pageIndex", "1"))
  if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, errors.New("integer parsing error (pageIndex)"))
    return
  }
  if pageIndex < 1 {
		httputil.NewError(ctx, http.StatusBadRequest, errors.New("pageIndex < 1"))
    return
  }
	pageLength, err = strconv.Atoi(ctx.DefaultQuery("pageLength", "1"))
  if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, errors.New("integer parsing error (pageLength)"))
    return
  }
  if pageLength < 1 {
		httputil.NewError(ctx, http.StatusBadRequest, errors.New("pageLength < 1"))
    return
  }
	result, err := getAttempts(ctx, true)
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}
	tweets := result.Data
	util.Reverse(&tweets)
	usersById := make(map[string]twitter.Profile, len(result.Includes.Users))
	for _, user := range result.Includes.Users {
		usersById[user.ID] = user
	}
	n := len(tweets)
	from := util.Max(0, util.Min(pageLength*(pageIndex-1), n-1))
	res := make([]model.Tweet, util.Min(from+pageLength, n)-from)
	for i := range res {
		tweet := tweets[from+i]
		res[i] = model.MakeTweet(tweet, usersById[tweet.AuthorId])
	}
	ctx.JSON(http.StatusOK, model.Page[model.Tweet]{Entries: res, NumberOfPages: (n + pageLength - 1) / pageLength})
}

func getAttemptsStats(ctx *gin.Context) (model.Chart, error) {
	gameTracker, err := util.IdToObject(ctx, gameTrackersById)
	if err != nil {
		return nil, err
	}
	result, err := getAttempts(ctx, false)
	if err != nil {
		return nil, err
	}
	tweets := result.Data
	solution, err := gameTracker.Solution()
	if err != nil {
		return nil, err
	}
	chartAsMap := make(map[string]int)
	for _, tweet := range tweets {
		text := strings.ToLower(tweet.Text)
		var attempt string
		if strings.Contains(text, solution) {
			attempt = solution
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

// gameAttempts godoc
// @Summary     Retrieve game's attempts' frequencies
// @Tags        tvgames
// @Produce     json
// @Param       id path string true "Game to query"
// @Success     200 {object} model.Chart
// @Failure     400	{object}	model.Error "integer parsing error (id)"
// @Failure     404	{object}	model.Error "game id not found"
// @Router      /tvgames/{id}/attempts/stats [get]
func gameAttemptsStats(ctx *gin.Context) {
	chart, err := getAttemptsStats(ctx)
	if err == nil {
		ctx.JSON(http.StatusOK, chart)
  }
}

// gameAttempts godoc
// @Summary     Retrieve game's number of successes and failures
// @Tags        tvgames
// @Produce     json
// @Param       id path string true "Game to query"
// @Success     200 {object} model.BooleanChart
// @Failure     400	{object}	model.Error "integer parsing error (id)"
// @Failure     404	{object}	model.Error "game id not found"
// @Router      /tvgames/{id}/results [get]
func gameResults(ctx *gin.Context) {
	gameTracker, err := util.IdToObject(ctx, gameTrackersById)
	if err == nil {
    var result *twitter.ProfileTweets
		result, err = getAttempts(ctx, false)
		if err == nil {
			tweets := result.Data
      var solution string
			solution, err = gameTracker.Solution()
			if err == nil {
				successes := 0
				for _, tweet := range tweets {
					if strings.Contains(strings.ToLower(tweet.Text), solution) {
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

func init() {
	games = make([]model.Game, len(gameTrackers))
	i := 0
	for k, v := range gameTrackers {
		gameTrackersById[k] = &v
		games[i] = v.Game
		gamesById[k] = &v.Game
		i += 1
	}
}
