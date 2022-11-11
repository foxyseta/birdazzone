package tvgames

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

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
// @Success     200 {object} model.Game
// @Param       id	path	int	true	"ID to search"
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
// @Failure     404	{string}	string  "game id not found"
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
			fmt.Errorf("Missing solution getter for %T", gameTracker))
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
	return twitter.GetTweetsFromHashtag(query, util.LastGameDate(time.Now()))
}

// gameAttempts godoc
// @Summary     Retrieve game's attempts
// @Tags        tvgames
// @Produce     json
// @Param       id path string true "Game to query"
// @Param       pageIndex query int false "Number of the page to query" minimum(1) default(1)
// @Param       pageLength query int false "Length of the page to query" minimum(1) default(10)
// @Success     200 {object} model.Page[model.Tweet]
// @Failure     400	{string}	string  "integer parsing error (pageIndex)"
// @Failure     400	{string}	string  "pageIndex < 1"
// @Failure     400	{string}	string  "integer parsing error (pageLength)"
// @Failure     400	{string}	string  "pageLength < 1"
// @Failure     404	{string}	string  "game id not found"
// @Router      /tvgames/{id}/attempts [get]
func gameAttempts(ctx *gin.Context) {
	var pageIndex, pageLength int
	pageIndex, err := strconv.Atoi(ctx.DefaultQuery("pageIndex", "1"))
	pageLength, err = strconv.Atoi(ctx.DefaultQuery("pageLength", "1"))
	if err != nil {
		return
	}
	result, err := getAttempts(ctx, true)
	if err == nil {
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
	httputil.NewError(ctx, http.StatusInternalServerError, err)
}

// gameAttempts godoc
// @Summary     Retrieve game's attempts' frequencies
// @Tags        tvgames
// @Produce     json
// @Param       id path string true "Game to query"
// @Success     200 {object} model.Chart
// @Failure     404	{string}	string  "game id not found"
// @Router      /tvgames/{id}/attempts/stats [get]
func gameAttemptsStats(ctx *gin.Context) {
	gameTracker, err := util.IdToObject(ctx, gameTrackersById)
	if err != nil {
		return
	}
	tweets, err := twitter.GetTweetsFromHashtag(gameTracker.Game.Hashtag, "?max_results=20&exclude=replies")
	print(tweets)
	if err == nil {
		ctx.JSON(http.StatusOK, model.Page[model.Tweet]{NumberOfPages: 1})
	} else {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
	}
}

// gameAttempts godoc
// @Summary     Retrieve game's number of successes and failures
// @Tags        tvgames
// @Produce     json
// @Param       id path string true "Game to query"
// @Success     200 {object} model.BooleanChart
// @Failure     404	{string}	string  "game id not found"
// @Router      /tvgames/{id}/results [get]
func gameResults(ctx *gin.Context) {
	gameTracker, err := util.IdToObject(ctx, gameTrackersById)
	if err != nil {
		return
	}
	tweets, err := twitter.GetTweetsFromHashtag(gameTracker.Game.Hashtag, "?max_results=20&exclude=replies")
	print(tweets)
	if err == nil {
		ctx.JSON(http.StatusOK, model.Page[model.Tweet]{NumberOfPages: 1})
	} else {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
	}
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
