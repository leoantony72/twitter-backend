package repositories

import (
	"encoding/json"
	"fmt"

	"github.com/leoantony72/twitter-backend/timeline/internal/model"
)

func (t *TimelineRepo) GetTimeline(username string, start int) ([]model.Tweets, error) {
	Timeline_tweets := []model.Tweets{}
	// Temp_tweets := []model.Tweets{}
	Temp_tweet := model.Tweets{}
	stop := start + 10
	key := fmt.Sprintf(Redis_Timeline_key, username)
	tweets, err := t.redis.LRange(ctx, key, int64(start), int64(stop)).Result()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for k, Tweet := range tweets {
		fmt.Println(k)
		err := json.Unmarshal([]byte(Tweet), &Temp_tweet)
		if err != nil {
			return nil, err
		}

		Timeline_tweets = append(Timeline_tweets, Temp_tweet)
	}
	user := model.User_followers{}
	following := TempUserFollows{}

	following_rows, _ := t.db.Model(&user).Select("followee").Where("follower=?", username).Rows()

	for following_rows.Next() {
		t.db.ScanRows(following_rows, &following.Followings)
	}

	tweetkey := "users:" + "%s" + ":tweets"
	for _, user := range following.Followings {
		key := fmt.Sprintf(tweetkey, user)
		
		// t.redis.ZRange(ctx, key, 0, 3)
		result, _ := t.redis.ZRevRangeWithScores(ctx, key, 0, 3).Result()
		for _, tweet := range result {
			TestTweet := model.Tweets{}
			json.Unmarshal([]byte(tweet.Member.(string)), &TestTweet)
			Timeline_tweets = append(Timeline_tweets, TestTweet)
		}

	}
	// fmt.Println(tweets)
	return Timeline_tweets, nil
}
