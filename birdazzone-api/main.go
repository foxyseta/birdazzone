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

func runServer() {
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

	//swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(address())
}

func main() {
  runServer()
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
