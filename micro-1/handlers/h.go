package handlers

import "github.com/gin-gonic/gin"

func Ping(c *gin.Context) {
	c.String(200, "Hello microservice health check v1!")
}
