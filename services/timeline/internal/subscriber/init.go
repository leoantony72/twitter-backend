package subscriber

import (
	"github.com/leoantony72/twitter-backend/timeline/internal/ports"
	"github.com/streadway/amqp"
)

type TimelineSubscriber struct {
	ch   *amqp.Channel
	repo ports.TimelineRepository
}

func NewTimelineSubscriber(ch *amqp.Channel, repo ports.TimelineRepository) TimelineSubscriber {
	return TimelineSubscriber{
		ch:   ch,
		repo: repo,
	}
}
