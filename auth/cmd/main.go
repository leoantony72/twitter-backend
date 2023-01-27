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
	r := gin.Default()
	//Database connection
	db := database.StartPostgres()
	//give database to repository
	repo := repositories.NewUserPostgresRepo(db)

	//give repo to service
	service := services.NewUseCase(repo)

	// pkg.Keys()

	middleware := middleware.NewUserMiddleware(service)

	handler.NewUserHandler(service, r, middleware)

	err := pkg.RegisterService()
	if err != nil {
		return
	}
	r.Run(":8080")
}
