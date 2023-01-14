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
		auth.POST("/signup", handler.CreateUser)
		auth.POST("/login", handler.LoginUser)
		auth.POST("/logout", m.Authorization(), handler.LogoutUser)
		protect := auth.Group("/admin").Use(m.Authorization())
		{
			protect.GET("/", handler.Test)
		}
	}
	user := r.Group("/user")
	{
		user.GET("/", handler.GetById)
		user.DELETE("/", m.Authorization(), handler.DeleteUser)
		user.PUT("/", m.Authorization(), handler.UpdateUser)
		user.POST("/follow", m.Authorization(), handler.Follow)
		user.POST("/unfollow", m.Authorization(), handler.Unfollow)
	}

	return handler
}
