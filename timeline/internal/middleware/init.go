package middleware

import "github.com/leoantony72/twitter-backend/timeline/internal/ports"

type TimelineMiddleware struct {
	timelineService ports.TimelineService
}

func NewTimelineMiddleware(s ports.TimelineService) *TimelineMiddleware {
	return &TimelineMiddleware{
		timelineService: s,
	}
}
