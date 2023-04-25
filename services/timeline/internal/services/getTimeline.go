package services

import "github.com/leoantony72/twitter-backend/timeline/internal/model"

func (t *TimelineService) GetTimeline(username string, start int) ([]model.Tweets, error) {
	tweets, err := t.repo.GetTimeline(username, start)
	if err != nil {
		return nil, err
	}
	return tweets, nil
}
