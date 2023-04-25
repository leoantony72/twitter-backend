package repositories

import (
	"errors"
	"fmt"
	"time"

	"github.com/leoantony72/twitter-backend/tweet/internal/model"
	"github.com/redis/go-redis/v9"
)

type TempTweet struct {
	Likes    []string `redis:"likes"`
	Retweets []string `redis:"retweets"`
}

func (t *TweetRepo) GetTweetById(id string) (*model.Tweets, error) {
	like := model.Like{}
	retweet := model.Retweet{}
	tempTweet := TempTweet{}
	tweet := model.Tweets{}
	redis_key := "tweets:" + id
	redis_tweet_like_key := "tweets:" + id + ":like"
	redis_tweet_retweet_key := "tweets:" + id + ":retweet"
	t.redis.ZRange(ctx, redis_tweet_like_key, 0, 5).ScanSlice(&tempTweet.Likes)
	t.redis.ZRange(ctx, redis_tweet_retweet_key, 0, 5).ScanSlice(&tempTweet.Retweets)

	primary_ok := DoesKeyExist(t, redis_key)
	like_ok := DoesKeyExist(t, redis_tweet_like_key)
	retweet_ok := DoesKeyExist(t, redis_tweet_retweet_key)
	err := t.redis.HGetAll(ctx, redis_key).Scan(&tweet)
	fmt.Println("cached data", primary_ok, like_ok, retweet_ok, err)
	if (primary_ok) && (like_ok) && (retweet_ok) && (err == nil) {
		fmt.Println("here")
		tweet.Likes = tempTweet.Likes
		tweet.Retweets = tempTweet.Retweets
		var tm time.Time
		tm.UnmarshalText([]byte(tweet.Encoded_date))
		tweet.CreatedAt = tm
		return &tweet, nil
	}
	fmt.Println("DATABASE HIT")
	result := t.db.Model(&tweet).Where("id=?", id).Scan(&tweet)
	if result.RowsAffected == 0 {
		return nil, errors.New("invalid Tweet ID")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	likes := t.db.Model(&like).Select("username").Where("tweet_id=?", id).Scan(&tempTweet.Likes)
	retweets := t.db.Model(&retweet).Select("username").Where("tweet_id=?", id).Scan(&tempTweet.Retweets)
	if likes.Error != nil {
		return nil, result.Error
	}
	if retweets.Error != nil {
		return nil, result.Error
	}
	tweet.Likes = tempTweet.Likes
	tweet.Retweets = tempTweet.Retweets

	for _, username := range tempTweet.Likes {
		t.redis.ZAdd(ctx, redis_tweet_like_key, redis.Z{Score: 0, Member: username})
	}
	for _, username := range tempTweet.Retweets {
		t.redis.ZAdd(ctx, redis_tweet_retweet_key, redis.Z{Score: 0, Member: username})
	}
	encoded_date, _ := tweet.CreatedAt.MarshalText()
	tweet.Encoded_date = string(encoded_date)
	t.redis.HSet(ctx, redis_key, &tweet)
	// t.redis.ExpireAt(ctx, redis_key, time.Now().Add(time.Second*20))
	return &tweet, nil
}
