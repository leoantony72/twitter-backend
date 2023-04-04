package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/leoantony72/twitter-backend/timeline/internal/middleware"
	"github.com/leoantony72/twitter-backend/timeline/internal/ports"
)

type TimelineHandler struct {
	timeline_service ports.TimelineService
}

func NewTimelineHandler(s ports.TimelineService, m *middleware.TimelineMiddleware, r *gin.Engine) *TimelineHandler {
	handler := &TimelineHandler{
		timeline_service: s,
	}
	//service health endpoint
	r.GET("/check", handler.Check)
	timeline := r.Group("/timeline")
	{

		timeline.GET("/home", m.Authorization(), handler.Timeline)
	}

	return handler
}
