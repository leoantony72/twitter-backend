package services

import "github.com/leoantony72/twitter-backend/timeline/internal/ports"

type TimelineService struct {
	repo ports.TimelineRepository
}

func NewTimelineService(repo ports.TimelineRepository) ports.TimelineService {
	return &TimelineService{
		repo: repo,
	}
}
