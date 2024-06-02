package websocket

type Rooms struct {
	Clients []*Client
}

func (rooms *Rooms) AddClient(client *Client) {
	rooms.Clients = append(rooms.Clients, client)
}

func (room *Rooms) GetClient() []Client {
	var cs []Client

	for _, client := range room.Clients {
		cs = append(cs, *client)
	}
	return cs
}

func (room *Rooms) Publish(msg []byte) {
	for _, client := range room.Clients {
		client.Send(msg)
	}
}
