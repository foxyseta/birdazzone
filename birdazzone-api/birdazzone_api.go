package main

import (
  "net/http"
  "os"

  "github.com/gin-gonic/gin"
)

const EnvHost = "HOST"
const EnvPort = "PORT"
const FallbackHost = "0.0.0.0"
const FallbackPort = "8080"

func lookupEnvWithFallback(key string, fallback string) (string) {
  val, ok := os.LookupEnv(key)
  if ok {
    return val
  }
  return fallback
}

func address() (string) {
  return lookupEnvWithFallback(EnvHost, FallbackHost) + ":" +
         lookupEnvWithFallback(EnvPort, FallbackPort)
}

func main() {
  r := gin.Default()
  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "Welcome to Birdazzone API!",
    })
  })
  r.GET("/hello", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "Hello, Bird!",
    })
  })
  r.Run(address())
}
