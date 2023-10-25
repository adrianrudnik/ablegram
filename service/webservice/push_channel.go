package webservice

import (
	"sync"
	"sync/atomic"
)

var historyIdGen atomic.Uint64

type PushChannel struct {
	clients      map[string]map[*PushClient]bool
	addClient    chan *PushClient
	removeClient chan *PushClient
	broadcast    chan interface{}
	history      sync.Map
}

func NewPushChannel(broadcastChan chan interface{}) *PushChannel {
	return &PushChannel{
		clients:      make(map[string]map[*PushClient]bool),
		addClient:    make(chan *PushClient),
		removeClient: make(chan *PushClient),
		broadcast:    broadcastChan,
	}
}

func (c *PushChannel) Run() {
	for {
		select {
		case client := <-c.addClient:
			c.AddClient(client)
		case client := <-c.removeClient:
			c.RemoveClient(client)
		case message := <-c.broadcast:
			c.Broadcast(message)
		}
	}
}

func (c *PushChannel) AddClient(client *PushClient) {
	reg := c.clients[client.ID]

	if reg == nil {
		reg = make(map[*PushClient]bool)
		c.clients[client.ID] = reg
	}

	c.clients[client.ID][client] = true

	Logger.Info().Str("id", client.ID).Msg("Websocket client registered")

	// Send over the channels history to the client, to get the frontend into the correct state
	c.history.Range(func(k, v interface{}) bool {
		client.tx <- v
		return true
	})
}

func (c *PushChannel) RemoveClient(client *PushClient) {
	if _, ok := c.clients[client.ID]; ok {
		delete(c.clients[client.ID], client)
		close(client.tx)

		Logger.Info().Str("id", client.ID).Msg("Websocket client unregistered")
	}
}

func (c *PushChannel) Broadcast(message interface{}) {
	c.history.Store(historyIdGen.Add(1), message)
}
