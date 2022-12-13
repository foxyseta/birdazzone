package fantacitorio

import (
	"fmt"
	"net/http"
	"strings"

	"git.hjkl.gq/team13/birdazzone-api/model"
	"git.hjkl.gq/team13/birdazzone-api/twitter"
	"github.com/gin-gonic/gin"
)

func FantacitorioGroup(group *gin.RouterGroup) {
	group.GET("/politicians", getPoliticians)
	group.GET("/teams", getTeams)
}

// getPoliticians godoc
// @Summary Get all Fantacitorio politicians and their scores
// @Tags    fantacitorio
// @Produce json
// @Success 200 {array} model.Politician
// @Router  /fantacitorio/politicians [get]
func getPoliticians(ctx *gin.Context) {
	politicians, err := getPoliticiansPoints()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	} else {
		ctx.JSON(http.StatusOK, politicians)
	}
}

// getTeams godoc
// @Summary Get Fantacitorio teams as tweets
// @Tags    fantacitorio
// @Produce json
// @Param   username query    string false "Optional Twitter username to search for"
// @Success 200      {array}  model.Tweet
// @Failure 400      {object} model.Error "Incorrect syntax for a username"
// @Failure 404      {object} model.Error "No team with such username"
// @Router  /fantacitorio/teams [get]
func getTeams(ctx *gin.Context) {
	// TODO: implement me (TG-177)
	username, hasName := ctx.GetQuery("username")
	search := "#fantacitorio has:media"
	if hasName {
		if username == "" || strings.Contains(username, " ") {
			ctx.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Message: "incorrect syntax for a username"})
			return
		}
		search += " from:" + username
	}
	result, err := twitter.GetManyRecentTweetsFromQuery(search, "", "")
	ret := []string{}
	for _, m := range result.Includes.Medias {
		if m.Height == 512 && m.Width == 1024 {
			ret = append(ret, m.URL)
		}
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	if result.Meta.ResultCount == 0 {
		ctx.JSON(http.StatusNotFound, model.Error{Code: http.StatusNotFound, Message: "no team with such username"})
		return
	}
	fmt.Println(result.Includes.Medias)
	ctx.JSON(http.StatusNotImplemented, ret)
}
