package chess

import (
	"net/http"
	"testing"

	"git.hjkl.gq/team13/birdazzone-api/util"
)

func TestGetTvGames(t *testing.T) {
	getChessMoves(util.GetTestingGinContext())
}
