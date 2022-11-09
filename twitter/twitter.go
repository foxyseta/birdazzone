package twitter

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TestTwitter godoc
// @Summary     Test Twitter API
// @Tags        test
// @Produce     json
// @Param       query	path	string	true	"Query to search"
// @Success     200
// @Failure     401	{string}	string  "Bearer Token Error"
// @Router      /twitter/{query} [get]
func TestTwitter(ctx *gin.Context) {
	q := ctx.Param("query")
	resp := getTweetsFromQuery(q)
	if resp.Body != nil {
		returnTweetQuery(ctx, resp)
	} else {
		returnTweetErr(ctx, resp)
	}
}

// GameSolution godoc
// @Summary     Retrieve game's solution
// @Tags        games
// @Produce     json
// @Param       game	path	string	true	"Game to search"
// @Success     200
// @Failure     404	{string}	string  "param GAME not found"
// @Router      /{game}/solution [get]
func GameSolution(ctx *gin.Context) {
	q := ctx.Param("game")
	fmt.Println(q)
	if q == "ghigliottina" {
		sol := getGhigliottinaSolution()
		if sol != "" && sol != "ERR_USER" && sol != "ERR_TWEETS" {
			ctx.JSON(http.StatusOK, gin.H{
				"solution": sol})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": sol})
		}
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "param GAME not found"})
	}
}

func InitAPI(v1 *gin.RouterGroup) {
	v1.GET("/:game/solution", GameSolution)
}
