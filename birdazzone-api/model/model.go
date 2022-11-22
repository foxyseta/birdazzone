package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"git.hjkl.gq/team13/birdazzone-api/twitter"
	"git.hjkl.gq/team13/birdazzone-api/util"
	"github.com/paulmach/go.geojson"
	"github.com/swaggo/swag/example/celler/httputil"
)

// @Description Object returned on failed requests
type Error httputil.HTTPError

// @Description Parameters to query a single page
type PageQuery struct {
	Length int `json:"username" minimum:"1" example:"5"`
	Index  int `json:"index" minimum:"1" example:"5"`
}

func (pg *PageQuery) String() string {
	if pg == nil {
		return util.NilRepresentation
	}
	return fmt.Sprintf("page #%d", pg.Index)
}

// @Description A subsequence of the search results
type Page[T any] struct {
	Entries       []T `json:"entries"`
	NumberOfPages int `json:"number_of_pages" minimum:"1"`
}

// @Description A pair made of one positive counter and one negative counter
type BooleanChart struct {
	Label     string `json:"string" example:"Votes"`
	Positives int    `json:"positives" minimum:"0" example:"209"`
	Negatives int    `json:"negatives" minimum:"0" example:"318"`
}

func (bc *BooleanChart) String() string {
	if bc == nil {
		return util.NilRepresentation
	}
	return fmt.Sprintf("%s: %d VS %d", bc.Label, bc.Positives, bc.Negatives)
}

// @Description A possible value inside a chart, alongside its absolute
// @Description frequency
type ChartEntry struct {
	Value             string `json:"value" example:"parola"`
	AbsoluteFrequency int    `json:"absolute_frequency" minimum:"0" example:"34"`
}

func (ce *ChartEntry) String() string {
	if ce == nil {
		return util.NilRepresentation
	}
	return fmt.Sprintf("(%s: %d)", ce.Value, ce.AbsoluteFrequency)
}

// @Description A chart as a sequence of entries
type Chart []ChartEntry

// @Description A game which can be observed
type Game struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Hashtag string `json:"hashtag"`
	Logo    string `json:"logo"`
}

func (g *Game) String() string {
	if g == nil {
		return util.NilRepresentation
	}
	return fmt.Sprintf("#%d (%s #%s)", g.Id, g.Name, g.Hashtag)
}

// @Description A possible solution for a Game
type GameKey struct {
	Key  string `json:"key" example:"parola"`
	Date string `json:"date" example:"2022-11-17" format:"date"`
}

func (gk *GameKey) String() string {
	if gk == nil {
		return util.NilRepresentation
	}
	return fmt.Sprintf("%s: '%s'", gk.Date, gk.Key)
}

// @Description A Twitter user
type User struct {
	Username        string `json:"username" example:"mariorossi"`
	Name            string `json:"name" example:"Mario Rossi"`
	ProfileImageUrl string `json:"profile_image_url"`
}

func MakeUser(user twitter.UIDLookup) User {
	return User{
		Username:        user.Data.Username,
		Name:            user.Data.Name,
		ProfileImageUrl: user.Data.ProfileImageUrl,
	}
}

func (u *User) String() string {
	if u == nil {
		return util.NilRepresentation
	}
	return fmt.Sprintf("%s (@%s)", u.Name, u.Username)
}

// @Description Useful metrics describing a single Tweet
type Metrics struct {
	LikeCount    int `json:"like_count" minimum:"0" example:"122"`
	ReplyCount   int `json:"reply_count" minimum:"0" example:"42"`
	RetweetCount int `json:"retweet_count" minimum:"0" example:"15"`
}

func (m *Metrics) String() string {
	if m == nil {
		return util.NilRepresentation
	}
	return fmt.Sprintf("(Likes: %d, Replies: %d, Retweets: %d)", m.LikeCount, m.ReplyCount, m.RetweetCount)
}

type openStreetMapCoordinates struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

// @Description Map coordinates.
type Coordinates struct {
	Longitude float64 `json:"longitude" minimum:"-180" maximum:"180" example:"-74.026675"`
	Latitude  float64 `json:"latitude" minimum:"-90" maximum:"90" example:"40.683935"`
}

func StringToCoordinates(s string) (Coordinates, error) {
	var result Coordinates
	if s == "" {
		return result, errors.New("User has no location")
	}
	data, err := util.GetRequest(
		"https://nominatim.openstreetmap.org/search",
		false,
		[]any{},
		util.Pair[string, string]{First: "q", Second: s},
		util.Pair[string, string]{First: "format", Second: "jsonv2"},
	)
	if err != nil {
		return result, err
	}
	var arr []openStreetMapCoordinates
	err = json.Unmarshal(data, &arr)
	if len(arr) == 0 {
		return result, fmt.Errorf("No coordinates found for %s", s)
	}
	if err != nil {
		return result, err
	}
	result.Longitude, err = strconv.ParseFloat(arr[0].Lon, 64)
	if err != nil {
		return result, err
	}
	result.Latitude, err = strconv.ParseFloat(arr[0].Lat, 64)
	if err != nil {
		return result, err
	}
	return result, nil
}

func MakeCoordinates(l *geojson.Geometry, p twitter.Profile) *Coordinates {
	if l == nil {
		result, err := StringToCoordinates(p.Location)
		if err != nil {
			println(err.Error())
			return nil
		}
		println(p.Location)
		return &result
	}
	if l.IsPoint() {
		println("Point")
		return &Coordinates{
			Longitude: l.Point[1],
			Latitude:  l.Point[0],
		}
	}
	if l.BoundingBox != nil {
		println("Bounding box")
		bb := l.BoundingBox
		n := len(bb) / 2
		return &Coordinates{
			Latitude:  (bb[0] + bb[n]) / 2,
			Longitude: (bb[1] + bb[n+1]) / 2,
		}
	}
	return nil
}

// @Description A post published on Twitter
type Tweet struct {
	Text        string       `json:"text" example:"Hello, world!"`
	Author      User         `json:"author"`
	CreatedAt   string       `json:"created_at" format:"date-time"`
	Metrics     Metrics      `json:"metrics"`
	Coordinates *Coordinates `json:"coordinates"`
}

func MakeTweet(tweet twitter.ProfileTweet, author twitter.Profile, location *geojson.Geometry) Tweet {
	coordinates := MakeCoordinates(location, author)
	return Tweet{
		Text: tweet.Text,
		Author: User{
			Username:        author.Username,
			Name:            author.Name,
			ProfileImageUrl: author.ProfileImageUrl,
		},
		CreatedAt: tweet.CreatedAt,
		Metrics: Metrics{
			LikeCount:    tweet.PublicMetrics.LikeCount,
			ReplyCount:   tweet.PublicMetrics.ReplyCount,
			RetweetCount: tweet.PublicMetrics.RetweetCount,
		},
		Coordinates: coordinates,
	}
}

func (t *Tweet) String() string {
	if t == nil {
		return util.NilRepresentation
	}
	return fmt.Sprintf("\"%s\"\n%s, %s\n%s", t.Text, &t.Author, t.CreatedAt, &t.Metrics)
}
