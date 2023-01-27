package handler

import "github.com/gin-gonic/gin"

func (t *UserHandler) Check(c *gin.Context) {
	c.Status(200)
}
