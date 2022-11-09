package model

type Page struct {
  Length int `json:"username" minimum:"1" example:"5"`
  Index int `json:"index" minimum:"1" example:"5"`
}

type Game struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type User struct {
  Username string `json:"username" example:"mariorossi"`
  Name string `json:"name" example:"Mario Rossi"`
  ProfileImageUrl string `json:"profile_image_url"`
}

type Metrics struct {
  LikeCount int `json:"like_count" minimum:"0" example:"122"`
  ReplyCount int `json:"reply_count" minimum:"0" example:"42"`
  RetweetCount int `json:"retweet_count" minimum:"0" example:"15"`
}

type Tweet struct {
  Text string `json:"text" example:"Hello, world!"`
  Author User `json:"author"`
  CreatedAt string `json:"created_at" format:"date-time"`
  Metrics Metrics `json:"metrics"`
}
