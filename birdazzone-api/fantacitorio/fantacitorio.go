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
// @Success 200 {array} model.Politicians
// @Router  /fantacitorio/politicians [get]
func getPoliticians(ctx *gin.Context) {
	// TODO: implement me (TG-170, TG-171, TG-172, TG-173)
	ctx.JSON(http.StatusNotImplemented, []model.Politician{})
}

// getTeams godoc
// @Summary Get Fantacitorio teams as tweets
// @Tags    tvgames
// @Produce json
// @Param   username  query     string true "Optional username to search for"
// @Success 200 {array} string
// @Router  /fantacitorio/teams [get]
func getTeams(ctx *gin.Context) {
	// TODO: implement me (TG-177)
	ctx.JSON(http.StatusNotImplemented, []model.Tweet{})
}
