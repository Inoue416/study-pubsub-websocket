package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

var rooms = Rooms{}

func ServeWs(c *gin.Context) error {
	topic := c.Param("topic") // Get Url Param
	fmt.Printf("\n\nTopic Param: %s\n\n", topic)
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Error: %+v\n", err)
	}

	defer ws.Close()

	client := &Client{
		Ws: ws,
	}

	rooms.AddSubscription(&Subscription{
		Topic:  topic,
		Client: client,
	})

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Error: %+v\n", err)
			break
		}

		rooms.Publish(msg, topic)
	}
	return nil
}
