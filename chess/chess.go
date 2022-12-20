package chess

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ChessGroup(group *gin.RouterGroup) {
	group.GET("/:user/:game/:turn", getChessMoves)
}

// getChessMoves godoc
// @Summary Get all suggested moves for a given #birdchess tweet
// @Tags    chess
// @Produce json
// @Param   user path     string      true "player 1's username"
// @Param   game path     string      true "game identifier, i.e. its starting instant" Format(date-time)
// @Param   turn path     int         true "second player's turn number"
// @Success 200  {string} string      "The second player's move"
// @Success 204  {string} string      "No one has played yet"
// @Failure 400  {object} model.Error "game > now or turn < 0"
// @Failure 404  {object} model.Error "No user or no post found"
// @Router  /chess/{user}/{game}/{turn} [get]
func getChessMoves(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, nil)
}
