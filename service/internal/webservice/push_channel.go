package webservice

import (
	"github.com/adrianrudnik/ablegram/internal/pusher"
	"github.com/samber/lo"
	"sync"
	"time"
)

type PushChannel struct {
	clients               map[*PushClient]bool
	clientsLock           sync.RWMutex
	addClient             chan *PushClient
	removeClient          chan *PushClient
	broadcast             chan interface{}
	history               []interface{}
	historyLock           sync.RWMutex
	triggerHistoryCleanup func()
}

func NewPushChannel(broadcastChan chan interface{}) *PushChannel {
	return &PushChannel{
		clients:      make(map[*PushClient]bool),
		addClient:    make(chan *PushClient),
		removeClient: make(chan *PushClient),
		broadcast:    broadcastChan,
		history:      make([]interface{}, 0, 500),
	}
}

func (c *PushChannel) Run() {
	c.StartHistoryCompactor()

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
	// Ensure the client is not already registered
	c.clientsLock.RLock()
	if _, ok := c.clients[client]; ok {
		Logger.Warn().Str("client", client.ID).Msg("Websocket client already registered")
		c.clientsLock.RUnlock()
		return
	}
	c.clientsLock.RUnlock()

	// Register the new client
	c.clientsLock.Lock()
	c.clients[client] = true
	c.clientsLock.Unlock()

	Logger.Info().Str("client", client.ID).Msg("Websocket client registered")

	// Send over the channels history to the client, to get the frontend into the correct state
	c.historyLock.RLock()
	for _, msg := range c.history {
		client.tx <- msg
	}
	count := len(c.history)
	c.historyLock.RUnlock()

	Logger.Info().Int("messages", count).Str("client", client.ID).Msg("Websocket client received history")

}

func (c *PushChannel) RemoveClient(client *PushClient) {
	c.clientsLock.Lock()
	if _, ok := c.clients[client]; ok {
		delete(c.clients, client)
		close(client.tx)

		Logger.Info().Str("id", client.ID).Msg("Websocket client unregistered")
	}
	c.clientsLock.Unlock()
}

func (c *PushChannel) Broadcast(message interface{}) {
	c.historyLock.Lock()
	c.history = append(c.history, message)
	c.historyLock.Unlock()

	// Distribute message to all connected clients
	c.clientsLock.RLock()
	clients := c.clients
	for client := range clients {
		select {
		case client.tx <- message:
		}
	}
	c.clientsLock.RUnlock()

	c.triggerHistoryCleanup()
}

func (c *PushChannel) StartHistoryCompactor() {
	// Establish a debounced history cleaner
	c.triggerHistoryCleanup, _ = lo.NewDebounce(250*time.Millisecond, func() {
		c.historyLock.Lock()

		startCount := len(c.history)

		// Make the latest messages to first ones in the slice
		lo.Reverse(c.history)

		// Only keep the newest processing update
		c.history = pusher.FilterAllExceptFirst(c.history, func(v interface{}) bool {
			_, ok := v.(*pusher.ProcessingStatusPush)
			return ok
		})

		// Only keep the newest metrics update
		c.history = pusher.FilterAllExceptFirst(c.history, func(v interface{}) bool {
			_, ok := v.(*pusher.MetricUpdatePush)
			return ok
		})

		// Filter all file updates, keep the newest one per file
		m := lo.Uniq(lo.FilterMap(c.history, func(v interface{}, _ int) (string, bool) {
			x, ok := v.(*pusher.FileStatusPush)
			if !ok {
				return "", false
			}

			return x.ID, true
		}))

		for _, hit := range m {
			c.history = pusher.FilterAllExceptFirst(c.history, func(v interface{}) bool {
				x, ok := v.(*pusher.FileStatusPush)
				return ok && x.ID == hit
			})
		}

		lo.Reverse(c.history)

		// Increase capacity again
		c.history = append(make([]any, 0, len(c.history)+500), c.history...)

		Logger.Debug().Int("before", startCount).Int("after", len(c.history)).Msg("Websocket history compacted")

		c.historyLock.Unlock()
	})
}
