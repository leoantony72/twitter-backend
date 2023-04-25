package repositories

import (
	"context"

	"github.com/leoantony72/twitter-backend/tweet/internal/ports"
	"github.com/redis/go-redis/v9"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

var ctx = context.Background()

type TweetRepo struct {
	db    *gorm.DB
	redis *redis.Client
	mq    *amqp.Channel
}

func NewTweetRepo(db *gorm.DB, redis *redis.Client, mq *amqp.Channel) ports.TweetRepository {
	return &TweetRepo{
		db:    db,
		redis: redis,
		mq: mq,
	}
}

func CheckErr(err error) error {
	if err != nil {
		return err
	}
	return nil
}

func DoesKeyExist(t *TweetRepo, key string) bool {
	exists := t.redis.Exists(ctx, key).Val()
	return exists != 0
}
