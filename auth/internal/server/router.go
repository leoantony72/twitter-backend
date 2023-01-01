package server

import (
	"github.com/gin-gonic/gin"
	"github.com/leoantony72/twitter-backend/auth/internal/controller"
)

func SetupRoutes(r *gin.Engine) {
	public := r.Group("/api")
	{
		public.GET("/", controller.Home)
	}
}
