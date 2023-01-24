package ports

import (
	"github.com/leoantony72/twitter-backend/tweet/internal/model"
)

type TweetRepository interface {
	GetTweetById(id string) (*model.Tweets, error)
	GetTweetByUser(username string) (*[]model.Tweets, error)
	CreateTweet(tweet model.Tweets) error
	DeleteTweet(id, user string) error

	LikeTweet(like model.Like) error
	DislikeTweet(like model.Like) error

	ReTweet(retweet model.Retweet) error
	DeleteReTweet(retweet model.Retweet) error

	TweetAuthor(id string) (*model.Tweets, error)
}
