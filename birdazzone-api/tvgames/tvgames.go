package tvgames

import (
  "fmt"
	"net/http"
	"strconv"

  "git.hjkl.gq/team13/birdazzone-api/tvgames/ghigliottina"
	"git.hjkl.gq/team13/birdazzone-api/model"
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

var games []*model.Game

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
  game, err := idToObject(ctx, gamesById)
  if err == nil {
	  ctx.JSON(http.StatusOK, game)
  }
}

// gameSolution godoc
// @Summary     Retrieve game's solution
// @Tags        games
// @Produce     json
// @Param       game	path	string	true	"Game to search"
// @Success     200
// @Failure     404	{string}	string  "param GAME not found"
// @Router      /{game}/solution [get]
func gameSolution(ctx *gin.Context) {
  gameTracker, err := idToObject(ctx, gameTrackersById)
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

func idToObject[T any](ctx *gin.Context, data map[int]T) (T, error) {
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

func init() {
  games = make([]*model.Game, len(gameTrackers))
  i := 0
  for k, v := range gameTrackersById {
    gameTrackersById[k] = v
    games[i] = &v.Game
    gamesById[k] = &v.Game
    i += 1
  }
}
