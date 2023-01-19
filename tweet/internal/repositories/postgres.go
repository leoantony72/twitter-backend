package repositories

import (
	"github.com/leoantony72/twitter-backend/tweet/internal/ports"
	"gorm.io/gorm"
)

type TweetRepo struct {
	db *gorm.DB
}

func NewTweetRepo(db *gorm.DB) ports.TweetRepository {
	return &TweetRepo{
		db:db,
	}
}
