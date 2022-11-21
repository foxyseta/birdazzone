package birdazzone

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"git.hjkl.gq/team13/birdazzone-api/model"
	"git.hjkl.gq/team13/birdazzone-api/tvgames/gametracker"
	"git.hjkl.gq/team13/birdazzone-api/twitter"
	"git.hjkl.gq/team13/birdazzone-api/util"
)

var birdazzoneTracker = gametracker.GameTracker{
	Game: model.Game{
		Id:      1,
		Name:    "Birdazzone",
		Hashtag: "#birdazzone",
		Logo:    "/public/birdazzone.png"},
	Query:    "#birdazzone -from:birdazzone -is:retweet",
	Solution: solution,
}

func GetBirdazzoneTracker() gametracker.GameTracker {
	return birdazzoneTracker
}

func solution(dt *time.Time) (model.GameKey, error) {
	start_time, end_time := "", ""
	if dt != nil {
		start_time = util.LastInstantAtGivenTime(*dt, 0)
		end_time = util.LastInstantAtGivenTime(dt.AddDate(0, 0, 1), 0)
		if end_time > util.DateToString(time.Now()) {
			end_time = ""
		}
	}
	tweets, err := twitter.GetRecentTweetsFromQuery("La soluzione alla partita di #birdazzone di oggi era:", start_time, end_time, 10)

	if err != nil {
		return model.GameKey{}, err
	}
	if tweets.Meta.ResultCount == 0 {
		return model.GameKey{}, errors.New("couldn't find Birdazzone solution")
	}
	m := regexp.MustCompile(`La soluzione alla partita di #birdazzone di oggi era:\s([A-Z]|[a-z])+`)
	a := strings.ToLower(strings.Trim(m.FindString(tweets.Data[0].Text), "La soluzione alla partita di #birdazzone di oggi era: "))
	if len(a) > 0 {
		return model.GameKey{
			Key:  a,
			Date: tweets.Data[0].CreatedAt,
		}, nil
	}
	return model.GameKey{}, errors.New("couldn't find Birdazzone solution")
}
