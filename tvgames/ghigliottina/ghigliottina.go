package ghigliottina

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

var ghigliottinaTracker = gametracker.GameTracker{
	Game:     model.Game{
    Id: 0,
    Name: "Ghigliottina",
    Hashtag: "#ghigliottina#leredita",
    Logo: "/public/eredita.png"},
	Query:    "(#ghigliottina OR #leredita) -from:quizzettone -is:retweet",
	Solution: solution,
}

func GetGhigliottinaTracker() gametracker.GameTracker {
	return ghigliottinaTracker
}

func solution() (string, error) {
	dt := time.Now()
	x := util.LastInstantAtGivenTime(dt, 16)
	tweets, err := twitter.GetRecentTweetsFromQuery("La #parola della #ghigliottina de #leredita di oggi", x, 10)

	if err != nil {
		return "", err
	}
	if tweets.Meta.ResultCount == 0 {
		return "", errors.New("couldn't find Ghigliottina solution")
	}
	m := regexp.MustCompile(`La #parola della #ghigliottina de #leredita di oggi è:\s([A-Z]|[a-z])+`)
	a := strings.ToLower(strings.Trim(m.FindString(tweets.Data[0].Text), "La #parola della #ghigliottina de #leredita di oggi è: "))
	if a == "" {
		return "", errors.New("couldn't find Ghigliottina solution")
	}
	return a, nil
}
