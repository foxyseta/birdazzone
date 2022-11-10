package model

import (
	"fmt"

	"git.hjkl.gq/team13/birdazzone-api/twitter"
)

const nilRepresentation = "<nil>"

// @Description Parameters to query a single page
type PageQuery struct {
	Length int `json:"username" minimum:"1" example:"5"`
	Index  int `json:"index" minimum:"1" example:"5"`
}

func (pg *PageQuery) String() string {
	if pg == nil {
		return nilRepresentation
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
	Positives int `json:"positives" minimum:"0" example:"209"`
	Negatives int `json:"negatives" minimum:"0" example:"318"`
}

func (bc *BooleanChart) String() string {
	if bc == nil {
		return nilRepresentation
	}
	return fmt.Sprintf("[%d VS %d]", bc.Positives, bc.Negatives)
}

// @Description A possible value inside a chart, alongside its absolute
// @Description frequency
type ChartEntry struct {
	Value             string `json:"value" example:"parola"`
	AbsoluteFrequency int    `json:"absolute_frequency" minimum:"0" example:"34"`
}

func (ce *ChartEntry) String() string {
	if ce == nil {
		return nilRepresentation
	}
	return fmt.Sprintf("(%s: %d)", ce.Value, ce.AbsoluteFrequency)
}

// @Description A chart as a sequence of entries
type Chart []ChartEntry

type Game struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Hashtag string `json:"hashtag"`
}

func (g *Game) String() string {
	if g == nil {
		return nilRepresentation
	}
	return fmt.Sprintf("#%d (%s #%s)", g.Id, g.Name, g.Hashtag)
}

// @Description A Twitter user
type User struct {
	Username        string `json:"username" example:"mariorossi"`
	Name            string `json:"name" example:"Mario Rossi"`
	ProfileImageUrl string `json:"profile_image_url"`
}

func makeUser(user twitter.UIDLookup) User {
	return User{
		Username:        user.Data.Username,
		Name:            user.Data.Name,
		ProfileImageUrl: user.Data.ProfileImageUrl,
	}
}

func (u *User) String() string {
	if u == nil {
		return nilRepresentation
	}
	return fmt.Sprintf("%s (%s)", u.Name, u.Username)
}

// @Description Useful metrics describing a single Tweet
type Metrics struct {
	LikeCount    int `json:"like_count" minimum:"0" example:"122"`
	ReplyCount   int `json:"reply_count" minimum:"0" example:"42"`
	RetweetCount int `json:"retweet_count" minimum:"0" example:"15"`
}

func (m *Metrics) String() string {
	if m == nil {
		return nilRepresentation
	}
	return fmt.Sprintf("(Likes: %d, Replies: %d, Retweets: %d)", m.LikeCount, m.ReplyCount, m.RetweetCount)
}

// @Description A post published on Twitter
type Tweet struct {
	Text      string  `json:"text" example:"Hello, world!"`
	Author    User    `json:"author"`
	CreatedAt string  `json:"created_at" format:"date-time"`
	Metrics   Metrics `json:"metrics"`
}

func (t *Tweet) String() string {
	if t == nil {
		return nilRepresentation
	}
	return fmt.Sprintf("\"%s\"\n%s, %s\n%s", t.Text, &t.Author, t.CreatedAt, &t.Metrics)
}
