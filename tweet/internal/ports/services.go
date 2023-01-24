package ports

import "github.com/leoantony72/twitter-backend/tweet/internal/model"

type TweetService interface {
	GetTweetById(id string) (*model.Tweets, error)
	GetTweetByUser(username string) ([]*model.Tweets, error)
	CreateTweet(tweet model.Tweets) error
	DeleteTweet(id string) error
	LikeTweet(id string) error
	DislikeTweet(id string) error
	ReTweet(id string) error
	DeleteReTweet(id string) error
}
