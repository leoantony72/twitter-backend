package ports

import "github.com/leoantony72/twitter-backend/timeline/internal/model"

type TimelineRepository interface {
	GetTimeline(username string, start int) ([]model.Tweets, error)
	GetUserTimeline(username string, start int) ([]model.Tweets, error)
	GetFollowers(username string) []string
	AddTimelineEntry(tweet model.Tweets, username string) error
}
