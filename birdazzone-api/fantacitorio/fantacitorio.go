package fantacitorio

import (
	"net/http"

	"git.hjkl.gq/team13/birdazzone-api/model"
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
// @Param   username query    string false "Optional username to search for"
// @Success 200      {array}  string
// @Failure 400      {object} model.Error "Incorrect syntax for a username"
// @Failure 404      {object} model.Error "No user with such username"
// @Router  /fantacitorio/teams [get]
func getTeams(ctx *gin.Context) {
	// TODO: implement me (TG-177)
	ctx.JSON(http.StatusNotImplemented, []model.Tweet{})
}
