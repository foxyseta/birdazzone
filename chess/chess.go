package chess

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"git.hjkl.gq/team13/birdazzone-api/model"
	"git.hjkl.gq/team13/birdazzone-api/twitter"
	"github.com/gin-gonic/gin"
)

func ChessGroup(group *gin.RouterGroup) {
	group.GET("/:user/:game/:turn", getChessMove)
}

// getChessMove godoc
// @Summary Get all suggested moves for a given #birdchess tweet
// @Tags    chess
// @Produce json
// @Param   user path     string      true "player 1's username"
// @Param   game path     string      true "game identifier, i.e. its starting instant" Format(date-time)
// @Param   turn path     int         true "second player's turn number"                minimum(1)
// @Success 200  {string} string      "The second player's move. It is recognized by the [a-h][1-8][a-h][1-8] regexp."
// @Success 204  {string} string      "No one has played yet"
// @Failure 400  {object} model.Error "username parsing module or integer parsing error on turn or turn <= 0"
// @Failure 404  {object} model.Error "No user or no post found"
// @Router  /chess/{user}/{game}/{turn} [get]
func getChessMove(ctx *gin.Context) {
	// Username
	username := ctx.Param("user")
	validUsername, _ := regexp.MatchString("^[A-Za-z0-9_]+$", username)
	if !validUsername {
		ctx.JSON(http.StatusBadRequest, model.Error{
			Code:    http.StatusBadRequest,
			Message: "Username parsing error",
		})
		return
	}
	// Date
	date := ctx.Param("game")
	minimumLimit := time.Now().AddDate(0, 0, -7).Format(time.RFC3339)
	if date < minimumLimit {
		date = minimumLimit
	}
	// Turn
	turn, err := strconv.Atoi(ctx.Param("turn"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Error{
			Code:    http.StatusBadRequest,
			Message: "Integer parsing error on 'turn' parameter",
		})
		return
	}
	if turn < 1 {
		ctx.JSON(http.StatusBadRequest, model.Error{
			Code:    http.StatusBadRequest,
			Message: "turn < 1",
		})
		return
	}
	// Result
	res, error := uncheckedGetCheckMove(username, date, turn)
	if error != nil {
		ctx.JSON(error.Code, model.Error{
			Code:    error.Code,
			Message: error.Message,
		})
		return
	}
	if res == "" {
		ctx.JSON(http.StatusNoContent, "")
		return
	}
	ctx.JSON(http.StatusOK, res)
}

const MaxResults = 100

func uncheckedGetCheckMove(username string, date string, turn int) (string, *model.Error) {
	tweets, err := twitter.GetRecentTweetsFromQuery("from:"+username, date, "", MaxResults)
	if err != nil {
		return "", &model.Error{Code: http.StatusNotFound, Message: err.Error()}
	}
	length := len(tweets.Data)
	if length < turn {
		return "", &model.Error{
			Code:    http.StatusNotFound,
			Message: fmt.Sprintf("Wanted to retrieve turn #%d. Tweets found: just %d.", turn, length),
		}
	}
	return mostPopularChessMove(tweets.Data[length-turn]), nil
}

var chessMove = regexp.MustCompile("[a-h][1-8][a-h][1-8]")

func mostPopularChessMove(tweet twitter.ProfileTweet) string {
	tweets, err := twitter.GetRecentQuotingTweets(tweet.ID)
	if err != nil {
		return ""
	}
	counters := map[string]int{}
	maxMove, maxCounter := "", 0
	for _, tweet := range tweets.Data {
		move := chessMove.FindString(strings.ToLower(tweet.Text))
		likes := 1 + tweet.PublicMetrics.LikeCount
		if move != "" {
			_, ok := counters[move]
			if ok {
				counters[move] += likes
			} else {
				counters[move] = likes
			}
			newCounter := counters[move]
			if newCounter > maxCounter {
				maxMove, maxCounter = move, newCounter
			}
		}
	}
	return maxMove
}
