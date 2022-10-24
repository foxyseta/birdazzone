// @title       Birdazzone API
// @version     0.1
// @description API implementation for the Software Engineering (90106) project for the academic year 2022/23 at the University of Bologna.

// @license.name GNU AFFERO GENERAL PUBLIC LICENSE
// @license.url  https://www.gnu.org/licenses/agpl-3.0.en.html

// @host     localhost:8080
// @BasePath /api/v1

package main

import (
	"fmt"
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

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")

	v1.GET("/hello", helloWorld)

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
	fmt.Println("HELLO")
}
