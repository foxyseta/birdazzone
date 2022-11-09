package tvgames

import "git.hjkl.gq/team13/birdazzone-api/model"

type GameTracker interface {
	// The game being tracked
	Game() model.Game
	// the hashtag used by the users playing this TV game
	Hashtags() string
	// the solution to the last instance of this TV game
	Solution() string
}
