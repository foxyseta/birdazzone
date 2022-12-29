package birdazzone

import (
	"errors"
	"strings"
	"time"

	"git.hjkl.gq/team13/birdazzone-api/model"
	"git.hjkl.gq/team13/birdazzone-api/tvgames/gametracker"
	"git.hjkl.gq/team13/birdazzone-api/twitter"
)

var birdazzoneTracker = gametracker.GameTracker{
	Game: model.Game{
		Name:    "Birdazzone",
		Hashtag: "#birdazzone",
		Logo:    "/public/birdazzone.png"},
	Query:        "#birdazzone -from:birdazzone -is:retweet",
	Solution:     givenSolution,
	LastSolution: lastSolution,
	Start:        4,
}

func GetBirdazzoneTracker() gametracker.GameTracker {
	return birdazzoneTracker
}

func solution(startTime string, endTime string) (model.GameKey, error) {
	tweets, err := twitter.GetRecentTweetsFromQuery("La soluzione al #birdazzone di oggi", startTime, endTime, 10)

	if err != nil {
		return model.GameKey{}, err
	}
	if tweets.Meta.ResultCount == 0 {
		return model.GameKey{}, errors.New("couldn't find Birdazzone solution")
	}
	text := tweets.Data[0].Text
	a := strings.ToLower(text[strings.LastIndex(text, " ")+1:])
	return gametracker.MakeGameKey("Birdazzone", a, tweets.Data[0].CreatedAt)
}

func givenSolution(dt time.Time) (model.GameKey, error) {
	return gametracker.GivenSolution(dt, solution)
}

func lastSolution() (model.GameKey, error) {
	return gametracker.LastSolution(solution)
}
