package websocket

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

var rooms = Rooms{}

func ServeWs(c *gin.Context) error {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Error: %+v\n", err)
	}

	defer ws.Close()

	client := &Client{
		Ws: ws,
	}

	rooms.AddClient(client)

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Error: %+v\n", err)
			break
		}

		rooms.Publish(msg)
	}
	return nil
}
