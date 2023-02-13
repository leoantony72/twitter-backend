package repositories

import (
	"context"

	"github.com/leoantony72/twitter-backend/timeline/internal/ports"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var ctx = context.Background()

type TimelineRepo struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewTimelineRepo(db *gorm.DB, redis *redis.Client) ports.TimelineRepository {
	return &TimelineRepo{
		db:    db,
		redis: redis,
	}
}
