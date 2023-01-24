package model

import "time"

type Tweets struct {
	Id           string    `json:"id" gorm:"primaryKey;type:text"`
	Username     string    `json:"user_id" gorm:"type:text REFERENCES users(username)"`
	Content      string    `json:"tweet_content" gorm:"type:varchar(250);not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"not null"`
	LikeCout     int       `json:"like_count" gorm:"type:int;default:0"`
	RetweetCount int       `json:"retweet_count" gorm:"type:int;default:0"`
}

type Retweet struct {
	TweetId  string `json:"tweet_id" gorm:"primaryKey; type:text REFERENCES tweets(id)"`
	Username string `json:"user_id" gorm:"primaryKey; type:text REFERENCES users(username)"`
}

type Like struct {
	TweetId  string `json:"tweet_id" gorm:"primaryKey; type:text REFERENCES tweets(id)"`
	Username string `json:"user_id" gorm:"primaryKey; type:text REFERENCES users(username)"`
}
