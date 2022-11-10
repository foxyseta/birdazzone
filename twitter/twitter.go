package twitter

import (
	"encoding/json"
  "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const BearerToken = "AAAAAAAAAAAAAAAAAAAAAE4higEAAAAAIAkazaLrT4LHjJx2XFPsdVzEPe8%3DE7HE9wBq5B5b0m4F8uGmcslObTmQFccb9gppULjUwTNBGj1Td3"

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

func getRequest(url string) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+BearerToken)
	resp, err := client.Do(req)
	return resp, err
}

func GetUser(username string) *UIDLookup {
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

func getTweets(templateUrl string, id string, params string) *ProfileTweets {
	id = url.QueryEscape(id)
	resp, err := getRequest(fmt.Sprintf(templateUrl, id, params))

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

func GetTweetsFromUser(id string, params string) *ProfileTweets {
	return getTweets("https://api.twitter.com/2/users/%s/tweets/%s", id, params)
}

func GetTweetsFromHashtag(id string, params string) *ProfileTweets {
	return getTweets("https://api.twitter.com/2/tweets/search/recent/%s", id, params)
}
