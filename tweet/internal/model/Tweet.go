package model

import "time"

type Tweets struct {
	Id           string    `json:"id" gorm:"primaryKey;type:text"`
	UserId       string    `json:"user_id" gorm:"type:text REFERENCES users(id)"`
	Content      string    `json:"tweet_content" gorm:"type:varchar(250);not null"`
	CreatedAt    time.Time `json:"created_at"`
	LikeCout     int       `json:"like_count" gorm:"type:int;default:0"`
	RetweetCount int       `json:"retweet_count" gorm:"type:int;default:0"`
}

type Retweet struct {
	TweetId string `json:"tweet_id" gorm:"primaryKey; type:text REFERENCES tweets(id)"`
	UserId  string `json:"user_id" gorm:"primaryKey; type:text REFERENCES users(id)"`
}

type Like struct {
	TweetId string `json:"tweet_id" gorm:"primaryKey; type:text REFERENCES tweets(id)"`
	UserId  string `json:"user_id" gorm:"primaryKey; type:text REFERENCES users(id)"`
}
