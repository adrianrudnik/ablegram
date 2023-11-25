package pusher

import (
	"github.com/adrianrudnik/ablegram/internal/config"
	"github.com/adrianrudnik/ablegram/internal/pushermsg"
	"github.com/adrianrudnik/ablegram/internal/workload"
	"github.com/samber/lo"
	"sync"
	"time"
)

type PushManager struct {
	config                *config.Config
	clients               map[*PushClient]bool
	clientsLock           sync.RWMutex
	addClient             chan *PushClient
	removeClient          chan *PushClient
	broadcast             <-chan workload.PushMessage
	history               []workload.PushMessage
	historyLock           sync.RWMutex
	triggerHistoryCleanup func()
}

func NewPushManager(conf *config.Config, pushChan <-chan workload.PushMessage) *PushManager {
	return &PushManager{
		config:       conf,
		clients:      make(map[*PushClient]bool),
		addClient:    make(chan *PushClient),
		removeClient: make(chan *PushClient),
		broadcast:    pushChan,
		history:      make([]interface{}, 0, 500),
	}
}

func (c *PushManager) Run() {
	c.StartHistoryCompactor()

	for {
		select {
		case client := <-c.addClient:
			c.AddClient(client)

			// Admins get the IP detail, though we do not want it for demo mode,
			// as anyone could become admin and there is no need to expose the real IP to other demo admins.
			ip := client.Conn.RemoteAddr().String()
			if c.config.Behaviour.DemoMode {
				ip = "127.0.0.128"
			}

			c.Broadcast(pushermsg.NewUserWelcomePush(client.ID, client.Role, client.DisplayName, ip))
			c.Broadcast(pushermsg.NewAboutYouPush(client.ID))

			// Send a current list of active users towards the newly connected one
			c.clientsLock.RLock()
			for kClient := range c.clients {
				kIp := client.Conn.RemoteAddr().String()
				if c.config.Behaviour.DemoMode {
					kIp = "127.0.0.129"
				}

				c.Broadcast(pushermsg.NewUserCurrentPush(client.ID, kClient.ID, kClient.Role, kClient.DisplayName, kIp))
			}
			c.clientsLock.RUnlock()

		case client := <-c.removeClient:
			c.Broadcast(pushermsg.NewUserGoodbyePush(client.ID))
			c.RemoveClient(client)

		case message := <-c.broadcast:
			c.Broadcast(message)
		}
	}
}

func (c *PushManager) AddClient(client *PushClient) {
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
		if !CanClientReceive(msg, client) {
			continue
		}

		if v, ok := msg.(FilteredMessage); ok && client.Role == GuestRole {
			msg = v.FilteredVariant()
		}

		client.Tx <- msg
	}

	count := len(c.history)
	c.historyLock.RUnlock()

	Logger.Info().Int("messages", count).Str("client", client.ID).Msg("Websocket client received history")
}

func (c *PushManager) RemoveClient(client *PushClient) {
	c.clientsLock.Lock()
	if _, ok := c.clients[client]; ok {
		delete(c.clients, client)
		close(client.Tx)

		Logger.Info().Str("id", client.ID).Msg("Websocket client unregistered")
	}
	c.clientsLock.Unlock()
}

func (c *PushManager) Broadcast(message interface{}) {
	// Append the message to the history, it the interface tells us to, or if the interface is missing
	record := true // we keep everything that has no details about a specific behaviour
	if v, ok := message.(RecordMessage); ok {
		record = v.KeepInHistory()
	}

	if record {
		c.historyLock.Lock()
		c.history = append(c.history, message)
		c.historyLock.Unlock()
	}

	// Distribute message to all connected clients
	c.clientsLock.RLock()
	for client := range c.clients {
		if !CanClientReceive(message, client) {
			continue
		}

		if v, ok := message.(FilteredMessage); ok && client.Role == GuestRole {
			client.Tx <- v.FilteredVariant()
		} else {
			client.Tx <- message
		}
	}
	c.clientsLock.RUnlock()

	if record {
		c.triggerHistoryCleanup()
	}
}

func (c *PushManager) StartHistoryCompactor() {
	// Establish a debounced history cleaner
	c.triggerHistoryCleanup, _ = lo.NewDebounce(250*time.Millisecond, func() {
		c.historyLock.Lock()

		startCount := len(c.history)

		// Make the latest messages to first ones in the slice
		lo.Reverse(c.history)

		// Only keep the newest processing update
		c.history = FilterAllExceptFirst(c.history, func(v interface{}) bool {
			_, ok := v.(*pushermsg.ProcessingStatusPush)
			return ok
		})

		// Only keep the newest tag update
		c.history = FilterAllExceptFirst(c.history, func(v interface{}) bool {
			_, ok := v.(*pushermsg.TagUpdatePush)
			return ok
		})

		// Only keep the newest metrics update
		c.history = FilterAllExceptFirst(c.history, func(v interface{}) bool {
			_, ok := v.(*pushermsg.MetricUpdatePush)
			return ok
		})

		// Filter all file updates, keep the newest one per file
		m := lo.Uniq(lo.FilterMap(c.history, func(v interface{}, _ int) (string, bool) {
			x, ok := v.(*pushermsg.FileStatusPush)
			if !ok {
				return "", false
			}

			return x.ID, true
		}))

		for _, hit := range m {
			c.history = FilterAllExceptFirst(c.history, func(v interface{}) bool {
				x, ok := v.(*pushermsg.FileStatusPush)
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
