package main

import (
	"github.com/Inoue416/study-pubsub-websocket.git/websocket" // 追加
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
