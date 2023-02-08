package repositories

import (
	"errors"
	"fmt"
	"time"

	"github.com/leoantony72/twitter-backend/tweet/internal/model"
)

func (t *TweetRepo) GetTweetByUser(username string) (*[]model.Tweets, error) {
	tempTweet := model.Tweets{}
	tweet := []model.Tweets{}
	tweet_ids := []string{}

	redis_user_key := "users:" + username + ":tweets"
	t.redis.ZRange(ctx, redis_user_key, 0, 10).ScanSlice(&tweet_ids)
	var err error
	var tm time.Time
	for _, id := range tweet_ids {
		redis_tweet_key := "tweets:" + id
		err = t.redis.HGetAll(ctx, redis_tweet_key).Scan(&tempTweet)
		tm.UnmarshalText([]byte(tempTweet.Encoded_date))
		tempTweet.CreatedAt = tm
		tweet = append(tweet, tempTweet)
	}
	if err == nil {
		fmt.Println("Cached Data: ")
		return &tweet, nil
	}

	result := t.db.Model(&tweet).Select("id", "username", "created_at", "like_count", "retweet_count").Where("username=?", username).Scan(&tweet)
	if result.RowsAffected == 0 {
		return nil, errors.New("invalid Username")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &tweet, nil
}
