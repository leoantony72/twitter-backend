package ports

import (
	"github.com/leoantony72/twitter-backend/tweet/internal/model"
)

type TweetRepository interface{
    GetTweet(id string) (*model.Tweets,error)
    CreateTweet(*model.Tweets) error
    DeleteTweet(id string) error
    LikeTweet(id string) error
    DislikeTweet(id string) error
    ReTweet(id string) error
    DeleteReTweet(id string) error
}
