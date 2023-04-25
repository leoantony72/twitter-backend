package model

import "time"

type Tweets struct {
	Id           string    `json:"id" gorm:"primaryKey;type:text" redis:"id"`
	Username     string    `json:"user_id" gorm:"type:text REFERENCES users(username); constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" redis:"username"`
	Content      string    `json:"tweet_content" gorm:"type:varchar(250);not null" redis:"content"`
	CreatedAt    time.Time `json:"created_at" gorm:"not null" `
	Encoded_date string    `redis:"created_at" json:"-" gorm:"-"`
	LikeCount    int       `json:"like_count" gorm:"type:int;default:0" redis:"like_count"`
	RetweetCount int       `json:"retweet_count" gorm:"type:int;default:0" redis:"retweet_count"`
	Likes        []string  `json:"likes" redis:"-" gorm:"-"`
	Retweets     []string  `json:"retweets" redis:"-" gorm:"-"`
}