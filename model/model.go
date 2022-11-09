package model

type User struct {
  Username string `json:"username" example:"mariorossi"`
  Name string `json:"name" example:"Mario Rossi"`
  ProfileImageUrl string `json:"profile_image_url"`
}

type Metrics struct {
  LikeCount int `json:"like_count" example:"122"`
  ReplyCount int `json:"reply_count" example:"42"`
  RetweetCount int `json:"retweet_count" example:"15"`
}

type Tweet struct {
  Text string `json:"text" example:"Hello, world!"`
  Author User `json:"author"`
  CreatedAt string `json:"created_at" format:"date-time"`
  Metrics Metrics `json:"metrics"`
}
