package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leoantony72/twitter-backend/timeline/database"
	"github.com/leoantony72/twitter-backend/timeline/internal/handler"
	"github.com/leoantony72/twitter-backend/timeline/internal/middleware"
	"github.com/leoantony72/twitter-backend/timeline/internal/repositories"
	"github.com/leoantony72/twitter-backend/timeline/internal/services"
	"github.com/leoantony72/twitter-backend/timeline/pkg"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	db := database.StartPostgres()
	redis := database.StartRedis()
	repo := repositories.NewTimelineRepo(db, redis)
	service := services.NewTimelineService(repo)

	middleware := middleware.NewTimelineMiddleware(service)

	handler.NewTimelineHandler(service, middleware, r)

	err := pkg.RegisterService()
	if err != nil {
		return
	}
	PORT := pkg.GetEnv("PORT")
	r.Run(":" + PORT)
}
