package fantacitorio

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"git.hjkl.gq/team13/birdazzone-api/model"
	"git.hjkl.gq/team13/birdazzone-api/twitter"
	"github.com/gin-gonic/gin"
)

func FantacitorioGroup(group *gin.RouterGroup) {
	group.GET("/politicians", getPoliticians)
	group.GET("/teams", getTeams)
}

func extractPoliticianPoints(text string) []model.Politician {
	var ret []model.Politician
	lines := strings.Split(text, "\n")
	for i := 0; i < len(lines); i++ {
		re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`) //regex for points
		pointsStr := re.FindString(lines[i])
		lines[i] = strings.ToUpper(lines[i])

		if len(pointsStr) > 1 { // excludes all points made of 1 char (usually coming from links) or empty strings
			lines[i] = strings.ReplaceAll(lines[i], pointsStr, "")
			if strings.Contains(strings.ToUpper(lines[i]), "MALUS") {
				pointsStr = "-" + pointsStr
			}
			lines[i] = strings.ReplaceAll(lines[i], " PUNTI", "")
			lines[i] = strings.ReplaceAll(lines[i], "MALUS DI ", "")
			lines[i] = strings.ReplaceAll(lines[i], " PER ", "")
			lines[i] = strings.ReplaceAll(lines[i], " - ", "")
			fmt.Println(pointsStr + "|" + lines[i])
			points, err := strconv.Atoi(pointsStr)
			if err == nil {
				ret = append(ret, model.Politician{Name: lines[i], Score: points})
			}
		}

	}
	return ret
}

func getPoliticiansPoints() ([]model.Politician, error) {
	result, err := twitter.GetManyRecentTweetsFromQuery("from:fanta_citorio punti -squadre -squadra", "", "")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var arr []model.Politician
	tweets := result.Data
	for i := 0; i < result.Meta.ResultCount; i++ {
		arr = append(arr, extractPoliticianPoints(tweets[i].Text)...)
	}
	return arr, err
}

// getPoliticians godoc
// @Summary Get all Fantacitorio politicians and their scores
// @Tags    fantacitorio
// @Produce json
// @Success 200 {array} model.Politician
// @Router  /fantacitorio/politicians [get]
func getPoliticians(ctx *gin.Context) {
	// TODO: implement me (TG-170, TG-171, TG-172, TG-173)
	// LIST: https://docs.google.com/spreadsheets/d/1Y5oBN9omqOkV0WoGoLRYOgACOxKC5vksN0VoCXZ5ljY/edit#gid=1522144602
	politicians, err := getPoliticiansPoints()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, politicians)
}

// getTeams godoc
// @Summary Get Fantacitorio teams as tweets
// @Tags    fantacitorio
// @Produce json
// @Param   username query    string false "Optional Twitter username to search for"
// @Success 200      {array}  model.Tweet
// @Failure 400      {object} model.Error "Incorrect syntax for a username"
// @Failure 404      {object} model.Error "No user with such username"
// @Router  /fantacitorio/teams [get]
func getTeams(ctx *gin.Context) {
	// TODO: implement me (TG-177)
	username, hasName := ctx.GetQuery("username")
	search := "#fantacitorio has:media"
	if hasName {
		if username == "" {
			ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Message: "incorrect syntax for a username"})
			return
		}
		search += " from:" + username
	}
	result, err := twitter.GetManyRecentTweetsFromQuery(search, "", "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	if result.Meta.ResultCount == 0 {
		ctx.JSON(http.StatusNotFound, model.Error{Code: http.StatusNotFound, Message: "no user with such username"})
		return
	}
	fmt.Println(result)
	ctx.JSON(http.StatusNotImplemented, []model.Tweet{})
}
