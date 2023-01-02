package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leoantony72/twitter-backend/auth/database"
	"github.com/leoantony72/twitter-backend/auth/internal/handler"
	"github.com/leoantony72/twitter-backend/auth/internal/services"
	"github.com/leoantony72/twitter-backend/auth/internal/repositories"
)

func main() {
	r := gin.Default()
	//Database connection
	db := database.StartPostgres()

	//give database to repository
	repo := repositories.NewUserPostgresRepo(db)

	//give repo to service
	service := services.NewUseCase(repo)

	handler.NewUserHandler(service, r)

	r.Run(":8080")
}
