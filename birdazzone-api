package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	twitter "github.com/g8rswimmer/go-twitter"
	"github.com/gin-gonic/gin"
)

const BearerToken = "AAAAAAAAAAAAAAAAAAAAAE4higEAAAAAIAkazaLrT4LHjJx2XFPsdVzEPe8%3DE7HE9wBq5B5b0m4F8uGmcslObTmQFccb9gppULjUwTNBGj1Td3"

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
func returnTweetQuery(c *gin.Context, recentSearch *twitter.TweetRecentSearch) {
	send, _ := json.Marshal(recentSearch.LookUps)
	var v interface{}
	json.Unmarshal(send, &v)
	data := v.(map[string]interface{})

	c.JSON(http.StatusOK, gin.H{
		"tweets": data,
	})
}

// testTwitter godoc
// @Summary     Test Twitter API
// @Tags        test
// @Produce     json
// @Param       query	path	string	true	"Query to search"
// @Success     200
// @Router      /twitter/{query} [get]
func testTwitter(ctx *gin.Context) {
	q := ctx.Param("query")

	tweet := &twitter.Tweet{
		Authorizer: authorize{
			Token: BearerToken,
		},
		Client: http.DefaultClient,
		Host:   "https://api.twitter.com",
	}
	//optional fields for twitter search
	fieldOpts := twitter.TweetFieldOptions{
		TweetFields: []twitter.TweetField{twitter.TweetFieldCreatedAt, twitter.TweetFieldLanguage, twitter.TweetFieldGeo},
		PlaceFields: []twitter.PlaceField{twitter.PlaceFieldCountry},
	}
	//research options
	searchOpts := twitter.TweetRecentSearchOptions{
		MaxResult: 10,
	}

	recentSearch, err := tweet.RecentSearch(context.Background(), q, searchOpts, fieldOpts)
	var tweetErr *twitter.TweetErrorResponse
	switch {
	case errors.As(err, &tweetErr):
		printTweetError(tweetErr)
	case err != nil:
		fmt.Println(err)
	default:
		returnTweetQuery(ctx, recentSearch)
	}

}
