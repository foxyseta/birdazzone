package gametracker

import "git.hjkl.gq/team13/birdazzone-api/model"

type GameSolutionGetter func() (string, error)

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
