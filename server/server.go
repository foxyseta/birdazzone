package server

import (
	"os"

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

func corsMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	if c.Request != nil && c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}

func Address() string {
	return lookupEnvWithFallback(EnvHost, FallbackHost) + ":" +
		lookupEnvWithFallback(EnvPort, FallbackPort)
}

func CreateServer() *gin.Engine {
	r := gin.Default()
	r.Use(corsMiddleware)
	r.Static("/public/", "./public/")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
