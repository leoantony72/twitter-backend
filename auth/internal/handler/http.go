package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/leoantony72/twitter-backend/auth/internal/middleware"
	"github.com/leoantony72/twitter-backend/auth/internal/ports"
)

type UserHandler struct {
	userUseCase ports.UserUseCase
}

func NewUserHandler(u ports.UserUseCase, r *gin.Engine) *UserHandler {
	handler := &UserHandler{userUseCase: u}

	auth := r.Group("/auth")
	{
		auth.GET("/user", handler.GetById)
		auth.PUT("/user", handler.UpdateUser)
		auth.DELETE("/user", handler.DeleteUser)
		auth.POST("/signup", handler.CreateUser)
		auth.POST("/login", handler.LoginUser)
		protect := auth.Group("/admin").Use(middleware.Authorization())
		{
			protect.GET("/", handler.Test)
		}
	}

	return handler
}
