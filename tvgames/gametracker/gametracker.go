package gametracker

import (
	"fmt"
	"time"

	"git.hjkl.gq/team13/birdazzone-api/model"
	"git.hjkl.gq/team13/birdazzone-api/util"
)

type GameSolutionGetter func(time.Time) (model.GameKey, error)
type LastGameSolutionGetter func() (model.GameKey, error)

type GameTracker struct {
	Game         model.Game
	Query        string
	Solution     GameSolutionGetter
	LastSolution LastGameSolutionGetter
	Start        int
}

func (gt *GameTracker) String() string {
	if gt == nil {
		return "<nil>"
	}
	return gt.Query
}

func GivenSolution(dt time.Time, solution func(string, string) (model.GameKey, error)) (model.GameKey, error) {
	startTime := util.LastInstantAtGivenTime(dt, 0)
	endTime := util.LastInstantAtGivenTime(dt.AddDate(0, 0, 1), 0)
	if endTime > util.DateToString(time.Now()) {
		endTime = ""
	}
	return solution(startTime, endTime)
}

func LastSolution(solution func(string, string) (model.GameKey, error)) (model.GameKey, error) {
	return solution("", "")
}

func MakeGameKey(gameName string, gameAnswer string, date string) (model.GameKey, error) {
	if len(gameAnswer) > 0 {
		return model.GameKey{
			Key:  gameAnswer,
			Date: date,
		}, nil
	}
	return model.GameKey{}, fmt.Errorf("Couldn't find %s solution", gameName)
}
