package ghigliottina

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"git.hjkl.gq/team13/birdazzone-api/model"
	"git.hjkl.gq/team13/birdazzone-api/tvgames/gametracker"
	"git.hjkl.gq/team13/birdazzone-api/twitter"
)

var ghigliottinaTracker = gametracker.GameTracker{
	Game: model.Game{
		Name:    "Ghigliottina",
		Hashtag: "#ghigliottina#leredita",
		Logo:    "/public/leredita.png"},
	Query:        "(#ghigliottina OR #leredita) -from:quizzettone -is:retweet",
	Solution:     givenSolution,
	LastSolution: lastSolution,
	Start:        18,
}

func GetGhigliottinaTracker() gametracker.GameTracker {
	return ghigliottinaTracker
}

func solution(startTime string, endTime string) (model.GameKey, error) {
	tweets, err := twitter.GetRecentTweetsFromQuery("La #parola della #ghigliottina de #leredita di oggi -is:retweet", startTime, endTime, 10)

	if err != nil {
		return model.GameKey{}, err
	}
	if tweets.Meta.ResultCount == 0 {
		return model.GameKey{}, errors.New("couldn't find Ghigliottina solution")
	}
	m := regexp.MustCompile(`La #parola della #ghigliottina de #leredita di oggi è:\s([A-Z]|[a-z])+`)
	a := strings.ToLower(strings.Trim(m.FindString(tweets.Data[0].Text), "La #parola della #ghigliottina de #leredita di oggi è: "))
	return gametracker.MakeGameKey("Ghigliottina", a, tweets.Data[0].CreatedAt)
}

func givenSolution(dt time.Time) (model.GameKey, error) {
	return gametracker.GivenSolution(dt, solution)
}

func lastSolution() (model.GameKey, error) {
	return gametracker.LastSolution(solution)
}
