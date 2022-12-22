package chess

import (
	"net/http"
	"strconv"

	"git.hjkl.gq/team13/birdazzone-api/model"
	// "git.hjkl.gq/team13/birdazzone-api/twitter"
	"github.com/gin-gonic/gin"
)

func ChessGroup(group *gin.RouterGroup) {
	group.GET("/:user/:game/:turn", getChessMove)
}

// getChessMove godoc
// @Summary Get all suggested moves for a given #birdchess tweet
// @Tags    chess
// @Produce json
// @Param   user path     string      true "player 1's username"
// @Param   game path     string      true "game identifier, i.e. its starting instant" Format(date-time)
// @Param   turn path     int         true "second player's turn number"                minimum(1)
// @Success 200  {string} string      "The second player's move"
// @Success 204  {string} string      "No one has played yet"
// @Failure 400  {object} model.Error "integer parsing error on turn or turn <= 0"
// @Failure 404  {object} model.Error "No user or no post found"
// @Router  /chess/{user}/{game}/{turn} [get]
func getChessMove(ctx *gin.Context) {
	username := ctx.Param("username")
	date := ctx.Param("game")
	turn, err := strconv.Atoi(ctx.Param("turn"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Error{
			Code:    http.StatusBadRequest,
			Message: "Integer parsing error on 'turn' parameter",
		})
		return
	}
	if turn < 1 {
		ctx.JSON(http.StatusBadRequest, model.Error{
			Code:    http.StatusBadRequest,
			Message: "turn < 1",
		})
		return
	}
	res, error := uncheckedGetCheckMove(username, date, turn)
	if error == nil {
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(error.Code, model.Error{
			Code:    error.Code,
			Message: error.Message,
		})
	}
}

func uncheckedGetCheckMove(username string, date string, turn int) (string, *model.Error) {
	return "e2e4", nil
}
