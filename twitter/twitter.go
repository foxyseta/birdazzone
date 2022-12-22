package twitter

import (
	"encoding/json"
	"fmt"
	"strconv"

	"git.hjkl.gq/team13/birdazzone-api/util"
	geojson "github.com/paulmach/go.geojson"
)

const BaseUrl = "https://api.twitter.com/2/"

type Profile struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Username        string `json:"username"`
	Location        string `json:"location"`
	ProfileImageUrl string `json:"profile_image_url"`
}

type Place struct {
	ID  string           `json:"id"`
	Geo geojson.Geometry `json:"geo"`
}

type Media struct {
	MediaKey string `json:"media_key"`
	Type     string `json:"type"`
	URL      string `json:"url"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
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
	Attachments struct {
		MediaKeys []string `json:"media_keys"`
	} `json:"attachments"`
	EditHistoryTweetIds []string `json:"edit_history_tweet_ids"`
	ID                  string   `json:"id"`
	Text                string   `json:"text"`
	Geo                 *struct {
		PlaceId string `json:"place_id"`
	} `json:"geo"`
	Coordinates *struct {
		Coordinates []float32 `json:"coordinates"`
	} `json:"coordinates"`
}

// List of tweets from a single profile
type ProfileTweets struct {
	Data     []ProfileTweet `json:"data"`
	Includes struct {
		Medias []Media   `json:"media"`
		Users  []Profile `json:"users"`
		Places []Place   `json:"places"`
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

func GetUser(username string) (*UIDLookup, error) {
	res, err := util.GetRequest("https://api.twitter.com/2/users/by/username/%s",
		true,
		[]any{username},
		util.Pair[string, string]{
			First:  "user.fields",
			Second: "id,name,username,location,profile_image_url",
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
	res, err := util.GetRequest(urlTemplate, true, pathParams, queryParams...)
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

func GetRecentTweetsFromQuery(query string, startTime string, endTime string, maxResults int) (*ProfileTweets, error) {
	return getTweets(
		"https://api.twitter.com/2/tweets/search/recent",
		[]any{},
		util.Pair[string, string]{First: "query", Second: query},
		util.Pair[string, string]{First: "start_time", Second: startTime},
		util.Pair[string, string]{First: "end_time", Second: endTime},
		util.Pair[string, string]{First: "max_results", Second: strconv.Itoa(maxResults)},
		util.Pair[string, string]{First: "tweet.fields", Second: "author_id,created_at,public_metrics,text,entities"},
		util.Pair[string, string]{First: "expansions", Second: "author_id,geo.place_id,attachments.media_keys"},
		util.Pair[string, string]{First: "user.fields", Second: "id,name,profile_image_url,username,location"},
		util.Pair[string, string]{First: "media.fields", Second: "preview_image_url,url,height,width"},
		util.Pair[string, string]{First: "place.fields", Second: "id,geo"},
	)
}

func GetRecentQuotingTweets(id string) (*ProfileTweets, error) {
	return getTweets(
		fmt.Sprintf("https://api.twitter.com/2/tweets/%s/quote_tweets", id),
		[]any{})
}

func GetManyRecentTweetsFromQuery(query string, startTime string, endTime string) (*ProfileTweets, error) {
	return GetRecentTweetsFromQuery(query, startTime, endTime, 100)
}
