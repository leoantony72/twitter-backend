package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leoantony72/twitter-backend/timeline/database"
	"github.com/leoantony72/twitter-backend/timeline/internal/handler"
	"github.com/leoantony72/twitter-backend/timeline/internal/middleware"
	"github.com/leoantony72/twitter-backend/timeline/internal/repositories"
	"github.com/leoantony72/twitter-backend/timeline/internal/services"
	"github.com/leoantony72/twitter-backend/timeline/internal/subscriber"
	"github.com/leoantony72/twitter-backend/timeline/pkg"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	db := database.StartPostgres()
	redis := database.StartRedis()
	mq := database.StartMQ()
	repo := repositories.NewTimelineRepo(db, redis)
	service := services.NewTimelineService(repo)
	subscriber.NewTimelineSubscriber(mq, repo)
	
	middleware := middleware.NewTimelineMiddleware(service)
	
	handler.NewTimelineHandler(service, middleware, r)
	
	err := pkg.RegisterService()
	if err != nil {
		return
	}
	go subscriber.ConsumeTweets(mq, repo)
	PORT := pkg.GetEnv("PORT")
	r.Run(":" + PORT)
}
