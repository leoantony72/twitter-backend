package ports

import "github.com/leoantony72/twitter-backend/timeline/internal/model"

type TimelineService interface {
	GetTimeline(username string, start int) ([]model.Tweets, error)
}
