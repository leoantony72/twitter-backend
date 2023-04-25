package repositories

import (
	"errors"
	"fmt"
	"time"

	"github.com/leoantony72/twitter-backend/tweet/internal/model"
	"github.com/redis/go-redis/v9"
)

func (t *TweetRepo) GetTweetByUser(username string) (*[]model.Tweets, error) {
	tempTweet := model.Tweets{}
	tweets := []model.Tweets{}
	tweet := model.Tweets{}
	tweet_ids := []string{}
	
	redis_user_key := "users:" + username + ":tweets"
	t.redis.ZRange(ctx, redis_user_key, 0, 10).ScanSlice(&tweet_ids)
	var err error
	var tm time.Time
	for _, id := range tweet_ids {
		// tweet_metadata_temp := TempTweet{}
		fmt.Println(id)
		redis_tweet_key := "tweets:" + id
		err = t.redis.HGetAll(ctx, redis_tweet_key).Scan(&tempTweet)
		tm.UnmarshalText([]byte(tempTweet.Encoded_date))
		tempTweet.CreatedAt = tm

		// redis_tweet_like_key := "tweets:" + id + ":like"
		// redis_tweet_retweet_key := "tweets:" + id + ":retweet"
		// t.redis.ZRange(ctx, redis_tweet_like_key, 0, 5).ScanSlice(&tweet_metadata_temp.Likes)
		// t.redis.ZRange(ctx, redis_tweet_retweet_key, 0, 5).ScanSlice(&tweet_metadata_temp.Retweets)
		// tempTweet.Likes = tweet_metadata_temp.Likes
		// tempTweet.Retweets = tweet_metadata_temp.Retweets
		tweets = append(tweets, tempTweet)
		fmt.Println("Cached Data fo sure")
	}
	if err == nil && len(tweets) != 0 {
		fmt.Println("Cached Data: ")
		return &tweets, nil
	}

	result := t.db.Model(&tweets).Select("id", "username", "content", "created_at", "like_count", "retweet_count").Where("username=?", username).Scan(&tweets)
	fmt.Println(tweet)
	if result.RowsAffected == 0 {
		return nil, errors.New("invalid Username or User has not tweeted")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	for _, tweet := range tweets {
		redis_tweet_key := "tweets:" + tweet.Id
		t.redis.ZAdd(ctx, redis_user_key, redis.Z{Score: 0, Member: tweet.Id})
		t.redis.HSet(ctx, redis_tweet_key, &tweet)
	}
	fmt.Println("DATABASE DATA")
	return &tweets, nil
}
