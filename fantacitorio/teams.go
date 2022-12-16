package fantacitorio

import (
	"net/http"
	"strings"

	"git.hjkl.gq/team13/birdazzone-api/model"
	"git.hjkl.gq/team13/birdazzone-api/twitter"
)

func teamsFromTwitter(username string, hasName bool) ([]model.FantaTeam, model.Error) {
	search := "#fantacitorio has:media -is:retweet"
	if hasName {
		if username == "" || strings.Contains(username, " ") {
			return []model.FantaTeam{}, model.Error{Code: http.StatusBadRequest, Message: "incorrect syntax for a username"}
		}
		search += " from:" + username
	}
	result, err := twitter.GetManyRecentTweetsFromQuery(search, "", "")
	if err != nil {
		return []model.FantaTeam{}, model.Error{Code: http.StatusInternalServerError, Message: err.Error()}
	}
	if result.Meta.ResultCount == 0 {
		return []model.FantaTeam{}, model.Error{Code: http.StatusNotFound, Message: "no team with such username"}
	}
	ret := tweetsToFantaTeam(result)
	return ret, model.Error{}
}

func tweetsToFantaTeam(tweets *twitter.ProfileTweets) []model.FantaTeam {
	var ret []model.FantaTeam
	for _, media := range tweets.Includes.Medias {
		if media.Width == 1024 && media.Height == 512 {
			ret = append(ret, mediaToFantaTeam(media, tweets)...)
		}
	}
	return ret
}

func mediaToFantaTeam(media twitter.Media, tweets *twitter.ProfileTweets) []model.FantaTeam {
	var ret []model.FantaTeam
	for _, data := range tweets.Data {
		if len(data.Attachments.MediaKeys) > 0 && media.MediaKey == data.Attachments.MediaKeys[0] {
			ret = append(ret, dataToFantaTeam(media, data, tweets)...)
			break
		}
	}
	return ret
}

func dataToFantaTeam(media twitter.Media, data twitter.ProfileTweet, tweets *twitter.ProfileTweets) []model.FantaTeam {
	var ret []model.FantaTeam
	for _, user := range tweets.Includes.Users {
		if data.AuthorId == user.ID {
			ret = append(ret,
				model.FantaTeam{
					Username:        user.Username,
					ProfileImageUrl: user.ProfileImageUrl,
					PostUrl:         "https://twitter.com/anyuser/status/" + data.ID,
					ImageUrl:        media.URL,
				})

			break
		}
	}
	return ret
}
