// @title       Birdazzone API
// @version     0.1
// @description API implementation for the Software Engineering (90106) project for the academic year 2022/23 at the University of Bologna.

// @license.name GNU AFFERO GENERAL PUBLIC LICENSE
// @license.url  https://www.gnu.org/licenses/agpl-3.0.en.html

// @host     localhost:8080
// @BasePath /api/v1

package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	_ "git.hjkl.gq/team13/birdazzone-api/docs"
	twitter "github.com/g8rswimmer/go-twitter"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const EnvHost = "HOST"
const EnvPort = "PORT"
const FallbackHost = "0.0.0.0"
const FallbackPort = "8080"

const BearerToken = "AAAAAAAAAAAAAAAAAAAAAE4higEAAAAAIAkazaLrT4LHjJx2XFPsdVzEPe8%3DE7HE9wBq5B5b0m4F8uGmcslObTmQFccb9gppULjUwTNBGj1Td3"

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
	v1 := r.Group("/api/v1")
	v1.GET("/hello", helloWorld)

	//test twitter API
	v1.GET("/test/:query", testTwitter)

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

// testTwitter godoc
// @Summary     Test Twitter API
// @Tags        test
// @Produce     json
// @Param       query	path	string	true	"Query to search"
// @Success     200
// @Router      /test/{query} [get]
func testTwitter(ctx *gin.Context) {
	query := ctx.Param("query")
	test(ctx, query)

}

//TWITTER TEST
func test(ctx *gin.Context, q string) {
	tweet := &twitter.Tweet{
		Authorizer: authorize{
			Token: BearerToken,
		},
		Client: http.DefaultClient,
		Host:   "https://api.twitter.com",
	}
	fieldOpts := twitter.TweetFieldOptions{
		TweetFields: []twitter.TweetField{twitter.TweetFieldCreatedAt, twitter.TweetFieldConversationID, twitter.TweetFieldLanguage},
	}
	searchOpts := twitter.TweetRecentSearchOptions{}

	recentSearch, err := tweet.RecentSearch(context.Background(), q, searchOpts, fieldOpts)
	var tweetErr *twitter.TweetErrorResponse
	switch {
	case errors.As(err, &tweetErr):
		printTweetError(tweetErr)
	case err != nil:
		fmt.Println(err)
	default:
		sendTweetQuery(ctx, recentSearch)
	}
}

type authorize struct {
	Token string
}

func (a authorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

func printTweetError(tweetErr *twitter.TweetErrorResponse) {
	enc, err := json.MarshalIndent(tweetErr, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(enc))
}
func sendTweetQuery(c *gin.Context, recentSearch *twitter.TweetRecentSearch) {
	send, _ := json.Marshal(recentSearch.LookUps)
	var v interface{}
	json.Unmarshal(send, &v)
	data := v.(map[string]interface{})

	c.JSON(http.StatusOK, data)
}
