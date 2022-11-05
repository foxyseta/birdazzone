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
	"os"
	"strconv"
	_ "git.hjkl.gq/team13/birdazzone-api/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const EnvHost = "HOST"
const EnvPort = "PORT"
const FallbackHost = "0.0.0.0"
const FallbackPort = "8080"

func lookupEnvWithFallback(key string, fallback string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	return fallback
}

func address() string {
	return lookupEnvWithFallback(EnvHost, FallbackHost) + ":" +
		lookupEnvWithFallback(EnvPort, FallbackPort)
}

func main() {
	//gin + API
	r := gin.Default()
  r.Use(corsMiddleware())
	v1 := r.Group("/api/v1")
	v1.GET("/hello", helloWorld)

	//twitter interaction API
	twttr := v1.Group("/twitter")
	twttr.GET("/ghigliottina", helloWorld) //TODO
	twttr.GET("/reazione", helloWorld)     //TODO
	twttr.GET("/:query", testTwitter)      //TODO: remove after /ghigliottina /reazione

  // Tv games
  games := v1.Group("/tv-games")
  games.GET("/", getTvGames)
  games.GET("/:id", getTvGameById)


	//swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(address())
}
func staticMessage(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}

// helloWorld godoc
// @Summary     Test your connection to birdazzone API
// @Tags        test
// @Produce     json
// @Success     200
// @Router      /hello [get]
func helloWorld(ctx *gin.Context) {
	staticMessage(ctx, "Hi! You've successfully connected to Birdazzone API.")

}

// TODO 
func corsMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

type Game struct {
  Id int `json:id`
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
  games := []Game {Game{Id: 0, Name: "La ghigliottina"}, Game{Id: 1, Name: "L'eredità"}}

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
  games := []Game {Game{Id: 0, Name: "La ghigliottina"}, Game{Id: 1, Name: "L'eredità"}}

  id := ctx.Param("id")
  num, err := strconv.Atoi(id)

  if err != nil || num >= len(games) || num < 0{
    ctx.AbortWithStatus(400)
    return 
  }

  ctx.JSON(http.StatusOK, games[num])
}

