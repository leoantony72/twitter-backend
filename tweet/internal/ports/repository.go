package ports

import (
	"github.com/leoantony72/twitter-backend/tweet/internal/model"
)

type TweetRepository interface {
	GetTweetById(id string) (*model.Tweets, error)
	GetTweetByUser(username string) ([]*model.Tweets, error)
	CreateTweet(tweet model.Tweets) error
	DeleteTweet(id,user string) error
	LikeTweet(id string) error
	DislikeTweet(id string) error
	ReTweet(id, user string) error
	DeleteReTweet(id, user string) error
}
