package fantacitorio

import (
	"net/http"
	"strings"

	"git.hjkl.gq/team13/birdazzone-api/model"
	"git.hjkl.gq/team13/birdazzone-api/twitter"
)

func teamsFromTwitter(username string, hasName bool) ([]string, model.Error) {
	ret := []string{}
	search := "#fantacitorio has:media"
	if hasName {
		if username == "" || strings.Contains(username, " ") {
			return ret, model.Error{Code: http.StatusBadRequest, Message: "incorrect syntax for a username"}
		}
		search += " from:" + username
	}
	result, err := twitter.GetManyRecentTweetsFromQuery(search, "", "")
	if err != nil {
		return ret, model.Error{Code: http.StatusInternalServerError, Message: err.Error()}
	}
	if result.Meta.ResultCount == 0 {
		return ret, model.Error{Code: http.StatusNotFound, Message: "no team with such username"}
	}
	for _, m := range result.Includes.Medias {
		if m.Height == 512 && m.Width == 1024 {
			ret = append(ret, m.URL)
		}
	}
	return ret, model.Error{}
}
