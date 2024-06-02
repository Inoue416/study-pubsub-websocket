package websocket

import "github.com/gorilla/websocket"

type Client struct {
	Ws *websocket.Conn
}

func (client *Client) Send(msg []byte) error {
	return client.Ws.WriteMessage(websocket.TextMessage, msg)
}
