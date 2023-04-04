package services

import (
	"github.com/leoantony72/twitter-backend/timeline/internal/ports"
)

type TimelineService struct {
	repo ports.TimelineRepository
}

const Redis_Timeline_key = "user:" + "%s" + ":timeline"

func NewTimelineService(repo ports.TimelineRepository) ports.TimelineService {
	return &TimelineService{
		repo: repo,
	}
}
