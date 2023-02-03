package repositories

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/leoantony72/twitter-backend/auth/internal/model"
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

func Cache(u *UserPostgresRepo, key string, cmd string, values model.User, field string) (*model.User, error) {
	if key == "" {
		return nil, errors.New("error no key provided")
	}
	exist := doesKeyExist(u, key)
	switch cmd {
	case "hget":
		{
			if !exist {
				return nil, errors.New("invalid key")
			}
			u.redis.HGet(ctx, key, field).Scan(&values)
			return &values, nil
		}
	case "hset":
		{
			u.redis.HSet(ctx, key, values)
			u.redis.ExpireAt(ctx, key, time.Now().Add(time.Second*10))
		}
	case "hgetall":
		{
			if !exist {
				return nil, errors.New("invalid key")
			}
			cmd := u.redis.HGetAll(ctx, key)
			// Verify key exists
			err := cmd.Err()
			if err == redis.Nil {
				log.Printf("Empty record")
			}

			if err != nil {
				log.Printf("Error reading from Redis")
				return nil, err
			}

			scanErr := cmd.Scan(&values)
			if scanErr != nil {
				log.Printf("Error Scanning Object: %v", scanErr)
				return nil, scanErr
			}
			return &values, nil
		}
	default:
		return nil, errors.New("invalid command")
	}
	return nil, nil
}

func doesKeyExist(u *UserPostgresRepo, key string) bool {
	exists := u.redis.Exists(ctx, key).Val()
	return exists != 0
}
