package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leoantony72/twitter-backend/tweet/database"
	"github.com/leoantony72/twitter-backend/tweet/internal/handler"
	"github.com/leoantony72/twitter-backend/tweet/internal/middleware"
	"github.com/leoantony72/twitter-backend/tweet/internal/repositories"
	"github.com/leoantony72/twitter-backend/tweet/internal/services"
	"github.com/leoantony72/twitter-backend/tweet/pkg"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	db := database.StartPostgres()
	redis := database.StartRedis()
	MQ := database.StartMQ()
	repo := repositories.NewTweetRepo(db, redis, MQ)
	service := services.NewTweetService(repo)

	middleware := middleware.NewTweetMiddleware(service)

	handler.NewTweetHandler(service, middleware, r)

	err := pkg.RegisterService()
	if err != nil {
		return
	}
	PORT := pkg.GetEnv("PORT")
	r.Run(":" + PORT)
}
