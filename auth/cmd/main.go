package main

import (
	// "time"

	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/leoantony72/twitter-backend/auth/database"
	"github.com/leoantony72/twitter-backend/auth/internal/handler"
	"github.com/leoantony72/twitter-backend/auth/internal/middleware"
	"github.com/leoantony72/twitter-backend/auth/internal/repositories"
	"github.com/leoantony72/twitter-backend/auth/internal/services"
	"github.com/leoantony72/twitter-backend/auth/pkg"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	//Database connection
	db := database.StartPostgres()
	redis := database.StartRedis()
	//give database to repository
	repo := repositories.NewUserPostgresRepo(db,redis)

	//give repo to service
	service := services.NewUseCase(repo)

	// pkg.Keys()

	middleware := middleware.NewUserMiddleware(service)

	handler.NewUserHandler(service, r, middleware)

	err := pkg.RegisterService()
	if err != nil {
		return
	}
	PORT := pkg.GetEnv("PORT")
	r.Run(":" + PORT)
}
