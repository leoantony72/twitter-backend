package repositories

import (
	"encoding/json"
	"fmt"

	"github.com/leoantony72/twitter-backend/timeline/internal/model"
)

func (t *TimelineRepo) GetUserTimeline(username string, start int) ([]model.Tweets, error) {
	tweets := []model.Tweets{}
	tweetkey := "users:" + "%s" + ":tweets"
	key := fmt.Sprintf(tweetkey, username)
	stop := start + 10
	result, _ := t.redis.ZRevRangeWithScores(ctx, key, int64(start), int64(stop)).Result()
	for _, tweet := range result {
		TempTweet := model.Tweets{}
		json.Unmarshal([]byte(tweet.Member.(string)), &TempTweet)
		tweets = append(tweets, TempTweet)
	}
	return tweets, nil
}
