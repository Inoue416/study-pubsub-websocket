package websocket

type Subscription struct {
	Topic  string
	Client *Client
}

type Rooms struct {
	Subscriptions []*Subscription
}

func (rooms *Rooms) AddSubscription(subscription *Subscription) {
	rooms.Subscriptions = append(rooms.Subscriptions, subscription)
}

func (room *Rooms) GetSubscription(topic string) []Subscription {
	var subs []Subscription

	for _, sub := range room.Subscriptions {
		if sub.Topic == topic {
			subs = append(subs, *sub)
		}
	}
	return subs
}

func (room *Rooms) Publish(msg []byte, topic string) {
	subs := room.GetSubscription(topic)

	for _, sub := range subs {
		err := sub.Client.Send(msg)
		if err != nil {
			return
		}
	}
}
