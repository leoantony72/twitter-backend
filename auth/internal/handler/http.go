package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/leoantony72/twitter-backend/auth/internal/ports"
)

type UserHandler struct {
	userUseCase ports.UserUseCase
}

func NewUserHandler(u ports.UserUseCase, r *gin.Engine) *UserHandler {
	handler := &UserHandler{userUseCase: u}

	r.GET("/user", handler.GetById)

	return handler
}


