package tvgames

import (
	"net/http"
	"strconv"

	"git.hjkl.gq/team13/birdazzone-api/model"
	"github.com/gin-gonic/gin"
)

func TvGamesGroup(group *gin.RouterGroup) {
	group.GET("/", getTvGames)
	group.GET("/:id", getTvGameById)
}

// getTvGames godoc
// @Summary     Get all TV games
// @Tags        tvgames
// @Produce     json
// @Success     200
// @Router      /tvgames	[get]
func getTvGames(ctx *gin.Context) {
	games := []model.Game{
		{Id: 0, Name: "La ghigliottina"},
		{Id: 1, Name: "L'eredità"},
	}

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
	games := []model.Game{
		{Id: 0, Name: "La ghigliottina"},
		{Id: 1, Name: "L'eredità"},
	}

	id := ctx.Param("id")
	num, err := strconv.Atoi(id)

	if err != nil || num >= len(games) || num < 0 {
		ctx.AbortWithStatus(400)
		return
	}

	ctx.JSON(http.StatusOK, games[num])
}
