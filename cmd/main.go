package main

import (
	"study-pubsub-websocket/websocket"

	"github.com/gin-gonic/gin"
)

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	r := gin.Default()
	r.GET("/healthcheck", healthCheck)
	r.GET("/ws/:topic", func(c *gin.Context) {
		if err := websocket.ServeWs(c); err != nil {
			c.AbortWithError(500, err)
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}
	})
	r.Static("ws-client", "web")
	r.Run()
}
