package services

import "github.com/leoantony72/twitter-backend/timeline/internal/model"

func (t *TimelineService) GetUserTimeline(username string, start int) ([]model.Tweets, error) {
	tweets, err := t.repo.GetUserTimeline(username, start)
	if err != nil {
		return nil, err
	}
	return tweets, nil
}
