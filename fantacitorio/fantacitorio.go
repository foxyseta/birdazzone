package fantacitorio

import (
	"net/http"

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
// @Success 200      {array}  model.FantaTeam
// @Failure 400      {object} model.Error "Incorrect syntax for a username"
// @Failure 404      {object} model.Error "No team with such username"
// @Router  /fantacitorio/teams [get]
func getTeams(ctx *gin.Context) {
	res, err := teamsFromTwitter(ctx.GetQuery("username"))
	if err.Message != "" {
		ctx.JSON(err.Code, gin.H{"code": err.Code, "message": err.Message})
		return
	}
	ctx.JSON(http.StatusOK, res)
}
