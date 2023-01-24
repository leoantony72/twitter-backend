package ports

import "github.com/leoantony72/twitter-backend/tweet/internal/model"

type TweetService interface {
	GetTweetById(id string) (*model.Tweets, error)
	GetTweetByUser(username string) (*[]model.Tweets, error)
	CreateTweet(tweet model.Tweets) error
	DeleteTweet(id, user string) error
	LikeTweet(id, user string) error
	DislikeTweet(id, user string) error
	ReTweet(id, user string) error
	DeleteReTweet(id, user string) error
	TweetAuthor(id string) (*model.Tweets, error)
}
