package twitter

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const BearerToken = "AAAAAAAAAAAAAAAAAAAAAE4higEAAAAAIAkazaLrT4LHjJx2XFPsdVzEPe8%3DE7HE9wBq5B5b0m4F8uGmcslObTmQFccb9gppULjUwTNBGj1Td3"

type QueryResponse struct {
	Body   *string
	Code   int
	Status string
}

// Basic user profile data
type UIDLookup struct {
	Data struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
	} `json:"data"`
}

// List of tweets from a single profile
type ProfileTweets struct {
	Data []struct {
		EditHistoryTweetIds []string `json:"edit_history_tweet_ids"`
		ID                  string   `json:"id"`
		Text                string   `json:"text"`
	} `json:"data"`
	Meta struct {
		NextToken   string `json:"next_token"`
		ResultCount int    `json:"result_count"`
		NewestID    string `json:"newest_id"`
		OldestID    string `json:"oldest_id"`
	} `json:"meta"`
}

func timeFormat(dt time.Time) string {
	// before 7PM returns yesterday's solution
	if dt.Hour() < 19 {
		dt = dt.AddDate(0, 0, -1)
	}
	// time format YYYY-MM-DDTHH:MM:SSZ (ISO 8601/RFC 3339 UTC TIMEZONE)
	x := strconv.Itoa(dt.Year()) + "-"
	if int(dt.Month()) < 10 {
		x += "0"
	}
	x += strconv.Itoa(int(dt.Month())) + "-"
	if dt.Day() < 10 {
		x += "0"
	}
	x += strconv.Itoa(dt.Day()) + "T18:55:00Z"
	return x
}

func getRequest(url string) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+BearerToken)
	resp, err := client.Do(req)
	return resp, err
}

func getTweetsFromQuery(query string) QueryResponse {
	max_results := "10"
	query = url.QueryEscape(query)
	resp, err := getRequest("https://api.twitter.com/2/tweets/search/recent?query=" + query + "&max_results=" + max_results + "&tweet.fields=public_metrics&expansions=geo.place_id&place.fields=geo,country,country_code&user.fields=username")

	if err != nil {
		log.Fatalln(err)
		return QueryResponse{nil, resp.StatusCode, resp.Status}
	}
	if resp.StatusCode != 200 {
		return QueryResponse{nil, resp.StatusCode, resp.Status}
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return QueryResponse{nil, resp.StatusCode, resp.Status}
	}
	//Convert the body to type string
	sb := string(body)

	return QueryResponse{&sb, resp.StatusCode, resp.Status}
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

func getTweetsFromUser(ID string, params string) *ProfileTweets {
	ID = url.QueryEscape(ID)
	resp, err := getRequest("https://api.twitter.com/2/users/" + ID + "/tweets" + params)

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
	var tweets ProfileTweets
	err = json.Unmarshal([]byte(string(body)), &tweets)
	if err != nil {
		panic(err)
	}

	return &tweets
}

func returnTweetQuery(c *gin.Context, res QueryResponse) {
	var v interface{}
	json.Unmarshal([]byte(*res.Body), &v)
	data := v.(map[string]interface{})
	c.JSON(http.StatusOK, data)
}

func returnTweetErr(c *gin.Context, res QueryResponse) {
	c.JSON(res.Code, res.Status)
}

// returns latest solution (lowercase word)
func getGhigliottinaSolution() string {
	user := getUser("quizzettone")
	if user == nil {
		return "ERR_USER"
	}
	dt := time.Now()
	x := timeFormat(dt)
	tweets := getTweetsFromUser(user.Data.ID, "?max_results=20&start_time="+x+"&exclude=replies")
	if tweets == nil {
		return "ERR_TWEETS"
	}
	for i := 0; i < tweets.Meta.ResultCount; i++ {
		if strings.Contains(tweets.Data[i].Text, "La #parola della #ghigliottina de #leredita di oggi è:") {
			m := regexp.MustCompile(`La #parola della #ghigliottina de #leredita di oggi è:\s([A-Z]|[a-z])+`)
			a := strings.ToLower(strings.Trim(m.FindString(tweets.Data[i].Text), "La #parola della #ghigliottina de #leredita di oggi è: "))
			return a
		}
	}
	return ""
}
