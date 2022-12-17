package chess

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ChessGroup(group *gin.RouterGroup) {
	group.GET("/:tweet", getChessMoves)
}

// getChessMoves godoc
// @Summary Get all suggested moves for a given #birdchess tweet
// @Tags    chess
// @Produce json
// @Param   tweet path     string      true "#birdchess tweet ID to query for"
// @Success 200   {array}  model.Chart "A chart whose values are the suggested moves"
// @Failure 404   {object} model.Error "No tweet with the given ID found"
// @Router  /chess/{tweet} [get]
func getChessMoves(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, nil)
}
