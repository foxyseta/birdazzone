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
	Game:     model.Game{Id: 0, Name: "Ghigliottina", Hashtag: "#ghigliottina#leredita"},
	Query:    "(#ghigliottina OR #leredita) -from:quizzettone -is:retweet",
	Solution: solution,
}

func GetGhigliottinaTracker() gametracker.GameTracker {
	return ghigliottinaTracker
}

func solution() (string, error) {
	user, err := twitter.GetUser("quizzettone")
	if err != nil {
		return "", err
	}
	dt := time.Now()
	x := util.LastInstantAtGivenTime(dt, 19)
	tweets, err := twitter.GetTweetsFromUser(user.Data.ID, 20, x)
	if err != nil {
		return "", err
	}
	for i := 0; i < tweets.Meta.ResultCount; i++ {
		if strings.Contains(tweets.Data[i].Text, "La #parola della #ghigliottina de #leredita di oggi è:") {
			m := regexp.MustCompile(`La #parola della #ghigliottina de #leredita di oggi è:\s([A-Z]|[a-z])+`)
			a := strings.ToLower(strings.Trim(m.FindString(tweets.Data[i].Text), "La #parola della #ghigliottina de #leredita di oggi è: "))
			return a, nil
		}
	}
	return "", errors.New("Couldn't find Ghigliottina solution")
}
