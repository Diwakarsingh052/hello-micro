package main

import (
	"github.com/gin-gonic/gin"
	"hello-service/handlers"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", handlers.Ping)
	r.GET("/hye/:name", func(c *gin.Context) {
		name := c.Param("name")
		if name == "" {
			c.String(http.StatusBadRequest, "Name is required")
			return
		}
		c.String(200, "Hello -----"+" "+name+"!")
	})
	log.Println("Listening on port 80")
	r.Run(":80")
}
