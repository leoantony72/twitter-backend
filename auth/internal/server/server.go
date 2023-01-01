package server

import "github.com/gin-gonic/gin"

func StartServer() {
	router := gin.Default()

	//Setup routes
	SetupRoutes(router)
	router.Run(":8080")

}
