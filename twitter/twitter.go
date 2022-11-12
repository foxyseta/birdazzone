package twitter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"git.hjkl.gq/team13/birdazzone-api/util"
)

const BaseUrl = "https://api.twitter.com/2/"
const BearerToken = "AAAAAAAAAAAAAAAAAAAAAE4higEAAAAAIAkazaLrT4LHjJx2XFPsdVzEPe8%3DE7HE9wBq5B5b0m4F8uGmcslObTmQFccb9gppULjUwTNBGj1Td3"

type Profile struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Username        string `json:"username"`
	ProfileImageUrl string `json:"profile_image_url"`
}

// Basic user profile data
type UIDLookup struct {
	Data Profile `json:"data"`
}

type ProfileTweet struct {
	AuthorId      string `json:"author_id"`
	CreatedAt     string `json:"created_at" format:"date-time"`
	PublicMetrics struct {
		LikeCount    int `json:"like_count"`
		ReplyCount   int `json:"reply_count"`
		RetweetCount int `json:"retweet_count"`
	} `json:"public_metrics"`
	EditHistoryTweetIds []string `json:"edit_history_tweet_ids"`
	ID                  string   `json:"id"`
	Text                string   `json:"text"`
}

// List of tweets from a single profile
type ProfileTweets struct {
	Data     []ProfileTweet `json:"data"`
	Includes struct {
		Users []Profile `json:"users"`
	} `json:"includes"`
	Meta struct {
		NextToken   string `json:"next_token"`
		ResultCount int    `json:"result_count"`
		NewestID    string `json:"newest_id"`
		OldestID    string `json:"oldest_id"`
	} `json:"meta"`
}

func toCompleteUrl(partialUrl string) string {
	return BaseUrl + partialUrl
}

func getRequest(
	urlTemplate string,
	pathParams []any,
	queryParams ...util.Pair[string, string],
) ([]byte, error) {
	if pathParams == nil {
		pathParams = []any{}
	}
	client := &http.Client{}

	// path parameters
	for i := range pathParams {
		pathParams[i] = url.PathEscape(fmt.Sprint(pathParams[i]))
	}
	req, err := http.NewRequest("GET", fmt.Sprintf(urlTemplate, pathParams...), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+BearerToken)

	// query parameters
	if queryParams != nil {
		q := req.URL.Query()
		for _, queryParam := range queryParams {
			q.Add(
				queryParam.First,
				queryParam.Second,
			)
		}
		req.URL.RawQuery = q.Encode()
	}
	println(req.URL.String())
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("From Twitter API: %s", resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return []byte(string(body)), nil
}

func GetUser(username string) (*UIDLookup, error) {
	res, err := getRequest("https://api.twitter.com/2/users/by/username/%s",
		[]any{username},
		util.Pair[string, string]{
			First:  "user.fields",
			Second: "id,name,username,profile_image_url",
		},
	)
	if err != nil {
		return nil, err
	}
	var uid UIDLookup
	err = json.Unmarshal(res, &uid)
	if err != nil {
		return nil, err
	}
	return &uid, nil
}

func getTweets(
	urlTemplate string,
	pathParams []any,
	queryParams ...util.Pair[string, string],
) (*ProfileTweets, error) {
	res, err := getRequest(urlTemplate, pathParams, queryParams...)
	if err != nil {
		return nil, err
	}
	var tweets ProfileTweets
	err = json.Unmarshal(res, &tweets)
	if err != nil {
		return nil, err
	}
	return &tweets, nil
}

func GetTweetsFromUser(id string, maxResults int, startTime string) (*ProfileTweets, error) {
	return getTweets(
		"https://api.twitter.com/2/users/%s/tweets",
		[]any{id},
		util.Pair[string, string]{First: "max_results", Second: strconv.Itoa(maxResults)},
		util.Pair[string, string]{First: "start_time", Second: startTime},
		util.Pair[string, string]{First: "exclude", Second: "replies"},
	)
}

func GetTweetsFromHashtag(query string, startTime string) (*ProfileTweets, error) {
	return getTweets(
		"https://api.twitter.com/2/tweets/search/recent",
		[]any{},
		util.Pair[string, string]{First: "query", Second: query},
		util.Pair[string, string]{First: "start_time", Second: startTime},
		util.Pair[string, string]{First: "max_results", Second: "100"},
		util.Pair[string, string]{First: "tweet.fields", Second: "author_id,created_at,public_metrics,text"},
		util.Pair[string, string]{First: "expansions", Second: "author_id"},
		util.Pair[string, string]{First: "user.fields", Second: "id,name,profile_image_url,username"},
	)
}
