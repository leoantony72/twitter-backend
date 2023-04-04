package repositories

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/leoantony72/twitter-backend/timeline/internal/model"
)

func (t *TimelineRepo) AddTimelineEntry(tweet model.Tweets, username string) error {
	key := fmt.Sprintf(Redis_Timeline_key, username)
	jsonData, err := json.Marshal(&tweet)
	if err != nil {
		return err
	}
	t.redis.LPush(ctx, key, jsonData)
	return nil
}
