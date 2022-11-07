package tvgames

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Game struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func InitAPI(v1 *gin.RouterGroup) {
	v1.GET("/tvgames", getTvGames)
	v1.GET("/tvgames/:id", getTvGameById)
}

// getTvGames godoc
// @Summary     Get all TV games
// @Tags        tvgames
// @Produce     json
// @Success     200
// @Router      /tvgames	[get]
func getTvGames(ctx *gin.Context) {
	// TODO
	games := []Game{{Id: 0, Name: "La ghigliottina"}, {Id: 1, Name: "L'eredità"}}

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
	// TODO
	games := []Game{{Id: 0, Name: "La ghigliottina"}, {Id: 1, Name: "L'eredità"}}

	id := ctx.Param("id")
	num, err := strconv.Atoi(id)

	if err != nil || num >= len(games) || num < 0 {
		ctx.AbortWithStatus(400)
		return
	}

	ctx.JSON(http.StatusOK, games[num])
}
