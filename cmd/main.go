package main

import (
	"study-websocket-pubsub/websocket"

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
	r.GET("/ws", websocket.ServeWs) // 追加
	r.Run()
}
