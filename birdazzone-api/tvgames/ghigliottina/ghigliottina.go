package ghigliottina

import (
	"regexp"
	"strings"
	"time"

	"git.hjkl.gq/team13/birdazzone-api/twitter"
	"git.hjkl.gq/team13/birdazzone-api/util"
)

func Solution() string {
	user := twitter.GetUser("quizzettone")
	if user == nil {
		return "ERR_USER"
	}
	dt := time.Now()
	x := util.TimeFormat(dt)
	tweets := twitter.GetTweetsFromUser(user.Data.ID, 20, x)
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
