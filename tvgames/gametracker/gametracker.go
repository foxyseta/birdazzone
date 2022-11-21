package gametracker

import (
	"time"

	"git.hjkl.gq/team13/birdazzone-api/model"
)

type GameSolutionGetter func(*time.Time) (model.GameKey, error)

type GameTracker struct {
	Game     model.Game
	Query    string
	Solution GameSolutionGetter
}

func (gt *GameTracker) String() string {
	if gt == nil {
		return "<nil>"
	}
	return gt.Query
}
