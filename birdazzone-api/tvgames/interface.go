package tvgames

type TvGame interface {
	// the hashtag used by the users playing this TV game
	Hashtags() string
	// the solution to the last instance of this TV game
	Solution() string
}
