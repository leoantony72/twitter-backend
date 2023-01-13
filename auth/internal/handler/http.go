package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/leoantony72/twitter-backend/auth/internal/middleware"
	"github.com/leoantony72/twitter-backend/auth/internal/ports"
)

type UserHandler struct {
	userUseCase ports.UserUseCase
}

func NewUserHandler(u ports.UserUseCase, r *gin.Engine, m *middleware.UserMiddleware) *UserHandler {
	handler := &UserHandler{userUseCase: u}

	auth := r.Group("/auth")
	{
		auth.GET("/user", handler.GetById)
		auth.PUT("/user", m.Authorization(), handler.UpdateUser)
		auth.DELETE("/user", m.Authorization(), handler.DeleteUser)
		auth.POST("/signup", handler.CreateUser)
		auth.POST("/login", handler.LoginUser)
		protect := auth.Group("/admin").Use(m.Authorization())
		{
			protect.GET("/", handler.Test)
		}
	}

	return handler
}
