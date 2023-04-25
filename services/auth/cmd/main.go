package main

import (
	// "time"

	// "github.com/gin-contrib/cors"
	"fmt"

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
	fmt.Println("Runnning 1")
	//Database connection
	db := database.StartPostgres()
	redis := database.StartRedis()
	//give database to repository
	repo := repositories.NewUserPostgresRepo(db, redis)
	fmt.Println("Runnning 2")

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
