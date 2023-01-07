package main

import (
	// "time"

	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/leoantony72/twitter-backend/auth/database"
	"github.com/leoantony72/twitter-backend/auth/internal/handler"
	"github.com/leoantony72/twitter-backend/auth/internal/repositories"
	"github.com/leoantony72/twitter-backend/auth/internal/services"
)

func main() {
	r := gin.Default()
	//Database connection
	db := database.StartPostgres()

	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"*"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	AllowOriginFunc: func(origin string) bool {
	// 		return origin == "https://github.com"
	// 	},
	// 	MaxAge: 12 * time.Hour,
	// }))
	//give database to repository
	repo := repositories.NewUserPostgresRepo(db)

	//give repo to service
	service := services.NewUseCase(repo)

	// pkg.Keys()

	handler.NewUserHandler(service, r)

	r.Run(":8080")
}
