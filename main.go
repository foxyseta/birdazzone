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

	_ "git.hjkl.gq/team13/birdazzone-api/docs"
	"git.hjkl.gq/team13/birdazzone-api/server"
	"git.hjkl.gq/team13/birdazzone-api/tvgames"
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
	tvgames.InitAPI(v1)

	server.Run(r)
}
