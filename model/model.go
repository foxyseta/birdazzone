package model

type PageQuery struct {
	Length int `json:"username" minimum:"1" example:"5"`
	Index  int `json:"index" minimum:"1" example:"5"`
}

type Page[T any] struct {
	Entries       []T `json:"entries"`
	NumberOfPages int `json:"number_of_pages" minimum:"1"`
}

type BooleanChart struct {
	Positives int `json:"positives" minimum:"0" example:"209"`
	Negatives int `json:"negatives" minimum:"0" example:"318"`
}

type ChartEntry struct {
	Value             string `json:"value" example:"parola"`
	AbsoluteFrequency int    `json:"absolute_frequency" minimum:"0" example:"34"`
}

type Game struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type User struct {
	Username        string `json:"username" example:"mariorossi"`
	Name            string `json:"name" example:"Mario Rossi"`
	ProfileImageUrl string `json:"profile_image_url"`
}

type Metrics struct {
	LikeCount    int `json:"like_count" minimum:"0" example:"122"`
	ReplyCount   int `json:"reply_count" minimum:"0" example:"42"`
	RetweetCount int `json:"retweet_count" minimum:"0" example:"15"`
}

type Tweet struct {
	Text      string  `json:"text" example:"Hello, world!"`
	Author    User    `json:"author"`
	CreatedAt string  `json:"created_at" format:"date-time"`
	Metrics   Metrics `json:"metrics"`
}
