package model

import (
  "fmt"

  "git.hjkl.gq/team13/birdazzone-api/twitter"
)

const nilRepresentation = "<nil>"

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

type Page[T any] struct {
	Entries       []*T `json:"entries"`
	NumberOfPages int `json:"number_of_pages" minimum:"1"`
}

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

type Chart []ChartEntry

type Game struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
  Hashtag string `json:"hashtag"`
}

func (g *Game) String() string {
  if g == nil {
    return nilRepresentation
  }
  return fmt.Sprintf("#%d (%s #%s)", g.Id, g.Name, g.Hashtag)
}

type User struct {
	Username        string `json:"username" example:"mariorossi"`
	Name            string `json:"name" example:"Mario Rossi"`
	ProfileImageUrl string `json:"profile_image_url"`
}

func makeUser(user twitter.UIDLookup) User {
  return User{Username: user.Data.Username, Name: user.Data.Name, ProfileImageUrl: ""}
}

func (u *User) String() string {
  if u == nil {
    return nilRepresentation
  }
  return fmt.Sprintf("%s (%s)", u.Name, u.Username)
}

type Metrics struct {
	LikeCount    int `json:"like_count" minimum:"0" example:"122"`
	ReplyCount   int `json:"reply_count" minimum:"0" example:"42"`
	RetweetCount int `json:"retweet_count" minimum:"0" example:"15"`
}

func (m *Metrics) String() string {
  if m == nil {
    return nilRepresentation
  }
  return fmt.Sprintf("(Likes: %d, Replies: %d, Retweets: %d)" , m.LikeCount, m.ReplyCount, m.RetweetCount)
}

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
  return fmt.Sprintf("\"%s\"\n%s, %s\n%s" , t.Text, &t.Author, t.CreatedAt, &t.Metrics)
}
