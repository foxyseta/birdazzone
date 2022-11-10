package tvgames

import (
  "fmt"
	"net/http"

  "git.hjkl.gq/team13/birdazzone-api/tvgames/ghigliottina"
	"git.hjkl.gq/team13/birdazzone-api/model"
	"git.hjkl.gq/team13/birdazzone-api/util"
	"github.com/gin-gonic/gin"
  "github.com/swaggo/swag/example/celler/httputil"
)

type GameSolutionGetter func() string

type GameTracker struct {
  Game model.Game
  Solution GameSolutionGetter
}

func (gt *GameTracker) String() string {
  if gt == nil {
    return "<nil>"
  }
  return gt.Game.String()
}

var gameTrackers = []GameTracker{
  { model.Game{Id: 0, Name: "Ghigliottina", Hashtag: "leredita"}, ghigliottina.Solution },
}

var gameTrackersById = map[int]*GameTracker{}

var games []model.Game

var gamesById = map[int]*model.Game{}

func TvGamesGroup(group *gin.RouterGroup) {
	group.GET("/", getTvGames)
	group.GET("/:id", getTvGameById)
	group.GET("/:id/solution", gameSolution)
}

// getTvGames godoc
// @Summary     Get all TV games
// @Tags        tvgames
// @Produce     json
// @Success     200
// @Router      /tvgames	[get]
func getTvGames(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, games)
}

// getTvGameById godoc
// @Summary     Get TV game
// @Tags        tvgames
// @Produce     json
// @Success     200
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
// @Param       game	path	string	true	"Game to search"
// @Success     200
// @Failure     404	{string}	string  "param GAME not found"
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
