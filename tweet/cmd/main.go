package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leoantony72/twitter-backend/tweet/database"
	"github.com/leoantony72/twitter-backend/tweet/internal/handler"
	"github.com/leoantony72/twitter-backend/tweet/internal/middleware"
	"github.com/leoantony72/twitter-backend/tweet/internal/repositories"
	"github.com/leoantony72/twitter-backend/tweet/internal/services"
)

func main() {
	r := gin.New()
	db := database.StartPostgres()
	repo := repositories.NewTweetRepo(db)
	service := services.NewTweetService(repo)

	middleware := middleware.NewTweetMiddleware(service)

	handler.NewTweetHandler(service, middleware, r)
	r.Run(":8090")
}
