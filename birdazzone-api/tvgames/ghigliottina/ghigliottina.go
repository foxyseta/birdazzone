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
	Game: model.Game{
		Id:      0,
		Name:    "Ghigliottina",
		Hashtag: "#ghigliottina#leredita",
		Logo:    "/public/leredita.png"},
	Query:    "(#ghigliottina OR #leredita) -from:quizzettone -is:retweet",
	Solution: solution,
}

func GetGhigliottinaTracker() gametracker.GameTracker {
	return ghigliottinaTracker
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
	tweets, err := twitter.GetRecentTweetsFromQuery("La #parola della #ghigliottina de #leredita di oggi", start_time, end_time, 10)

	if err != nil {
		return model.GameKey{}, err
	}
	if tweets.Meta.ResultCount == 0 {
		return model.GameKey{}, errors.New("couldn't find Ghigliottina solution")
	}
	m := regexp.MustCompile(`La #parola della #ghigliottina de #leredita di oggi è:\s([A-Z]|[a-z])+`)
	a := strings.ToLower(strings.Trim(m.FindString(tweets.Data[0].Text), "La #parola della #ghigliottina de #leredita di oggi è: "))
	if len(a) > 0 {
		return model.GameKey{
			Key:  a,
			Date: tweets.Data[0].CreatedAt,
		}, nil
	}
	return model.GameKey{}, errors.New("couldn't find Ghigliottina solution")
}
