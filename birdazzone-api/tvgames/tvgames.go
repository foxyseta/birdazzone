package tvgames

import (
	"fmt"
	"net/http"

	"git.hjkl.gq/team13/birdazzone-api/model"
	"git.hjkl.gq/team13/birdazzone-api/tvgames/ghigliottina"
	"git.hjkl.gq/team13/birdazzone-api/twitter"
	"git.hjkl.gq/team13/birdazzone-api/util"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
)

type GameSolutionGetter func() string

type GameTracker struct {
	Game     model.Game
	Solution GameSolutionGetter
}

func (gt *GameTracker) String() string {
	if gt == nil {
		return "<nil>"
	}
	return gt.Game.String()
}

var gameTrackers = []GameTracker{
	{model.Game{Id: 0, Name: "Ghigliottina", Hashtag: "leredita"}, ghigliottina.Solution},
}

var gameTrackersById = map[int]*GameTracker{}

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
		ctx.JSON(http.StatusOK, gin.H{
			"solution": gameTracker.Solution(),
		})
	} else {
		httputil.NewError(ctx, http.StatusInternalServerError,
			fmt.Errorf("Missing solution getter for %T", gameTracker))
	}
}

// gameAttempts godoc
// @Summary     Retrieve game's attempts
// @Tags        tvgames
// @Produce     json
// @Param       id path string true "Game to query"
// @Success     200 {object} model.Page[model.Tweet]
// @Failure     404	{string}	string  "game id not found"
// @Router      /tvgames/{id}/attempts [get]
func gameAttempts(ctx *gin.Context) {
	gameTracker, err := util.IdToObject(ctx, gameTrackersById)
	if err != nil {
		return
	}
	tweets := twitter.GetTweetsFromHashtag(gameTracker.Game.Hashtag, "?max_results=20&exclude=replies")
	if tweets != nil {
		ctx.JSON(http.StatusOK, model.Page[model.Tweet]{NumberOfPages: 1})
	} else {
		httputil.NewError(ctx, http.StatusInternalServerError,
			fmt.Errorf("Missing tweets"))
	}
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
	tweets := twitter.GetTweetsFromHashtag(gameTracker.Game.Hashtag, "?max_results=20&exclude=replies")
	if tweets != nil {
		ctx.JSON(http.StatusOK, model.Page[model.Tweet]{NumberOfPages: 1})
	} else {
		httputil.NewError(ctx, http.StatusInternalServerError,
			fmt.Errorf("Missing tweets"))
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
	tweets := twitter.GetTweetsFromHashtag(gameTracker.Game.Hashtag, "?max_results=20&exclude=replies")
	if tweets != nil {
		ctx.JSON(http.StatusOK, model.Page[model.Tweet]{NumberOfPages: 1})
	} else {
		httputil.NewError(ctx, http.StatusInternalServerError,
			fmt.Errorf("Missing tweets"))
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
