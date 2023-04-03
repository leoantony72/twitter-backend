package repositories

import (
	// "time"

	"encoding/json"

	"github.com/leoantony72/twitter-backend/tweet/internal/model"
	"github.com/leoantony72/twitter-backend/tweet/pkg"
	"github.com/redis/go-redis/v9"
	"github.com/streadway/amqp"
)

func (t *TweetRepo) CreateTweet(tweet model.Tweets) error {

	result := t.db.Create(&tweet)
	if result.Error != nil {
		return result.Error
	}
	redis_key := "tweets:" + tweet.Id
	user_redis_key := "users:" + tweet.Username + ":tweets"
	encoded_date, _ := tweet.CreatedAt.MarshalText()
	tweet.Encoded_date = string(encoded_date)
	t.redis.HSet(ctx, redis_key, &tweet)
	t.redis.ZAdd(ctx, user_redis_key, redis.Z{Score: 0, Member: tweet.Id})
	// t.redis.ExpireAt(ctx, redis_key, time.Now().Add(time.Second*20))

	data, err := json.Marshal(tweet)
	if err != nil {
		return err
	}

	ExchName := pkg.GetEnv("EXCHNAME")
	t.mq.Publish(ExchName, "", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(data),
	})
	return nil
}
