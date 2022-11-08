package twitter

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

const BearerToken = "AAAAAAAAAAAAAAAAAAAAAE4higEAAAAAIAkazaLrT4LHjJx2XFPsdVzEPe8%3DE7HE9wBq5B5b0m4F8uGmcslObTmQFccb9gppULjUwTNBGj1Td3"

type queryResponse struct {
	Body   *string
	Code   int
	Status string
}

type UIDLookup struct {
	Data struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
	} `json:"data"`
}

func getTweetsQuery(query string) queryResponse {
	max_results := "10"
	query = url.QueryEscape(query)
	resp, err := getRequest("https://api.twitter.com/2/tweets/search/recent?query=" + query + "&max_results=" + max_results + "&tweet.fields=public_metrics&expansions=geo.place_id&place.fields=geo,country,country_code&user.fields=username")

	if err != nil {
		log.Fatalln(err)
		return queryResponse{nil, resp.StatusCode, resp.Status}
	}
	if resp.StatusCode != 200 {
		return queryResponse{nil, resp.StatusCode, resp.Status}
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return queryResponse{nil, resp.StatusCode, resp.Status}
	}
	//Convert the body to type string
	sb := string(body)

	return queryResponse{&sb, resp.StatusCode, resp.Status}
}

func getUser(username string) *UIDLookup {

	username = url.QueryEscape(username)
	resp, err := getRequest("https://api.twitter.com/2/users/by/username/" + username)

	if err != nil {
		log.Fatalln(err)
		return nil
	}
	if resp.StatusCode != 200 {
		return nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	var uid UIDLookup
	err = json.Unmarshal([]byte(string(body)), &uid)
	if err != nil {
		panic(err)
	}
	return &uid
}

func getRequest(url string) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+BearerToken)
	resp, err := client.Do(req)
	return resp, err
}

func returnTweetQuery(c *gin.Context, res queryResponse) {
	var v interface{}
	json.Unmarshal([]byte(*res.Body), &v)
	data := v.(map[string]interface{})
	c.JSON(http.StatusOK, data)
}

func returnTweetErr(c *gin.Context, res queryResponse) {
	c.JSON(res.Code, res.Status)
}

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
	resp := getTweetsQuery(q)
	if resp.Body != nil {
		returnTweetQuery(ctx, resp)
	} else {
		returnTweetErr(ctx, resp)
	}
}
