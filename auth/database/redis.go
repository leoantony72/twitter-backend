package database

import "github.com/redis/go-redis/v9"

func StartRedis() *redis.Client {
	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return redis
}
