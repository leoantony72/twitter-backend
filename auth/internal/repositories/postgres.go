package repositories

import (
	"context"

	"github.com/leoantony72/twitter-backend/auth/internal/ports"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var ctx = context.Background()

type UserPostgresRepo struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewUserPostgresRepo(db *gorm.DB, redis *redis.Client) ports.UserRepository {
	return &UserPostgresRepo{
		db:    db,
		redis: redis,
	}
}

func DoesKeyExist(u *UserPostgresRepo, key string) bool {
	exists := u.redis.Exists(ctx, key).Val()
	return exists != 0
}
