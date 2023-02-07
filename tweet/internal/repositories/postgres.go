package repositories

import (
	"context"

	"github.com/leoantony72/twitter-backend/tweet/internal/ports"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var ctx =context.Background()

type TweetRepo struct {
	db *gorm.DB
	redis *redis.Client
}

func NewTweetRepo(db *gorm.DB, redis *redis.Client) ports.TweetRepository {
	return &TweetRepo{
		db:    db,
		redis: redis,
	}
}

func CheckErr(err error) error {
	if err != nil {
		return err
	}
	return nil
}
