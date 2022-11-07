// @title       Birdazzone API
// @version     0.1
// @description API implementation for the Software Engineering (90106) project for the academic year 2022/23 at the University of Bologna.

// @license.name GNU AFFERO GENERAL PUBLIC LICENSE
// @license.url  https://www.gnu.org/licenses/agpl-3.0.en.html

// @host     localhost:8080
// @BasePath /api/v1

package main

import (
	"net/http"
	"strconv"

	"git.hjkl.gq/team13/birdazzone-api/server"

	_ "git.hjkl.gq/team13/birdazzone-api/docs"
	"github.com/gin-gonic/gin"
)

// helloWorld godoc
// @Summary     Test your connection to birdazzone API
// @Tags        test
// @Produce     json
// @Success     200
// @Router      /hello [get]
func helloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hi! You've successfully connected to Birdazzone API.",
	})
}

func main() {
	r := server.CreateServer()
	v1 := r.Group("/api/v1")
	v1.GET("/hello", helloWorld)
	// Tv games
	v1.GET("/tv-games", getTvGames)
	v1.GET("/tv-games/:id", getTvGameById)
	server.Run(r)
}

type Game struct {
	Id   int    `json:id`
	Name string `json:name`
}

// getTvGames godoc
// @Summary     Get all TV games
// @Tags        tv-games
// @Produce     json
// @Success     200
// @Router      /tv-games
func getTvGames(ctx *gin.Context) {
	// TODO
	games := []Game{Game{Id: 0, Name: "La ghigliottina"}, Game{Id: 1, Name: "L'eredità"}}

	ctx.JSON(http.StatusOK, games)
}

// getTvGameById godoc
// @Summary     Get TV game
// @Tags        tv-games
// @Produce     json
// @Success     200
// @Router      /tv-games/{id} [get]
func getTvGameById(ctx *gin.Context) {
	// TODO
	games := []Game{Game{Id: 0, Name: "La ghigliottina"}, Game{Id: 1, Name: "L'eredità"}}

	id := ctx.Param("id")
	num, err := strconv.Atoi(id)

	if err != nil || num >= len(games) || num < 0 {
		ctx.AbortWithStatus(400)
		return
	}

	ctx.JSON(http.StatusOK, games[num])
}
