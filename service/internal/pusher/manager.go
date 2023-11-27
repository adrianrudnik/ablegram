package pusher

import (
	"github.com/adrianrudnik/ablegram/internal/auth"
	"github.com/adrianrudnik/ablegram/internal/config"
	"github.com/adrianrudnik/ablegram/internal/pushermsg"
	"github.com/adrianrudnik/ablegram/internal/workload"
	"github.com/samber/lo"
	"sync"
	"time"
)

type PushManager struct {
	config *config.Config

	clients     map[*Client]bool
	clientsLock sync.RWMutex

	users *auth.UserList

	addClient    chan *Client
	removeClient chan *Client
	broadcast    <-chan workload.PushMessage

	history               []workload.PushMessage
	historyLock           sync.RWMutex
	triggerHistoryCleanup func()
}

func NewPushManager(conf *config.Config, ul *auth.UserList, pushChan <-chan workload.PushMessage) *PushManager {
	return &PushManager{
		config:       conf,
		clients:      make(map[*Client]bool),
		users:        ul,
		addClient:    make(chan *Client),
		removeClient: make(chan *Client),
		broadcast:    pushChan,
		history:      make([]interface{}, 0, 500),
	}
}

func (pm *PushManager) Run() {
	pm.StartHistoryCompactor()

	for {
		select {
		case client := <-pm.addClient:
			pm.AddClient(client)

			// Admins get the IP detail, though we do not want it for demo mode,
			// as anyone could become admin and there is no need to expose the real IP to other demo admins.
			ip := client.GetIP(!pm.config.Behaviour.DemoMode)
			user := pm.users.Get(client.UserID)

			// Notify all clients about the new user
			pm.Broadcast(pushermsg.NewUserWelcomePush(
				ip, client.UserID, user.Role, user.DisplayName,
			))

			//// Notify the
			//c.Broadcast(pushermsg.NewAboutYouPush(client.ID))
			//

			// Send over a list of all currently connected clients
			for _, connClient := range pm.GetClients() {
				pm.Broadcast(pushermsg.NewUserCurrentPush(
					user.UserID,
					connClient.ID,
					connClient.GetIP(!pm.config.Behaviour.DemoMode),
					connClient.UserID,
					user.Role,
					user.DisplayName,
				))
			}

		case client := <-pm.removeClient:
			pm.Broadcast(pushermsg.NewUserGoodbyePush(client.ID))
			pm.RemoveClient(client)

		case message := <-pm.broadcast:
			pm.Broadcast(message)
		}
	}
}

func (pm *PushManager) AddClient(client *Client) {
	// Ensure the client is not already registered
	pm.clientsLock.RLock()
	if _, ok := pm.clients[client]; ok {
		Logger.Warn().Str("client", client.ID.String()).Msg("Websocket client already registered")
		pm.clientsLock.RUnlock()
		return
	}
	pm.clientsLock.RUnlock()

	// Register the new client
	pm.clientsLock.Lock()
	pm.clients[client] = true
	pm.clientsLock.Unlock()

	Logger.Info().Str("client", client.ID.String()).Msg("Websocket client registered")

	// Send over the channels history to the client, to get the frontend into the correct state
	pm.historyLock.RLock()
	for _, msg := range pm.history {
		if !pm.canClientReceive(msg, client) {
			continue
		}

		client.Tx <- pm.reduceMessageForClient(msg, client)
	}

	count := len(pm.history)
	pm.historyLock.RUnlock()

	Logger.Info().
		Int("messages", count).
		Str("client", client.ID.String()).
		Msg("Websocket client received history")
}

func (pm *PushManager) RemoveClient(client *Client) {
	pm.clientsLock.Lock()
	defer pm.clientsLock.Unlock()

	userId := client.UserID

	// Remove the client, end the send channel
	if _, ok := pm.clients[client]; ok {
		delete(pm.clients, client)
		close(client.Tx)

		Logger.Info().Str("id", client.ID.String()).Msg("Websocket client unregistered")
	}

	// If this was the last connected client with the given user ID, remove the user from the list
	// as he no longer has any active connections.
	wasLast := true
	for c, _ := range pm.clients {
		if c.UserID == userId {
			wasLast = false
		}
	}

	if wasLast {
		pm.users.Delete(userId)
	}
}

func (pm *PushManager) Broadcast(message interface{}) {
	// Append the message to the history, it the interface tells us to, or if the interface is missing
	record := true // we keep everything that has no details about a specific behaviour
	if v, ok := message.(RecordMessage); ok {
		record = v.KeepInHistory()
	}

	if record {
		pm.historyLock.Lock()
		pm.history = append(pm.history, message)
		pm.historyLock.Unlock()
	}

	// Distribute message to all connected clients
	pm.clientsLock.RLock()
	for client := range pm.clients {
		if !pm.canClientReceive(message, client) {
			continue
		}

		client.Tx <- pm.reduceMessageForClient(message, client)
	}
	pm.clientsLock.RUnlock()

	if record {
		pm.triggerHistoryCleanup()
	}
}

func (pm *PushManager) StartHistoryCompactor() {
	// Establish a debounced history cleaner
	pm.triggerHistoryCleanup, _ = lo.NewDebounce(250*time.Millisecond, func() {
		pm.historyLock.Lock()

		startCount := len(pm.history)

		// Make the latest messages to first ones in the slice
		lo.Reverse(pm.history)

		// Only keep the newest processing update
		pm.history = FilterAllExceptFirst(pm.history, func(v interface{}) bool {
			_, ok := v.(*pushermsg.ProcessingStatusPush)
			return ok
		})

		// Only keep the newest tag update
		pm.history = FilterAllExceptFirst(pm.history, func(v interface{}) bool {
			_, ok := v.(*pushermsg.TagUpdatePush)
			return ok
		})

		// Only keep the newest metrics update
		pm.history = FilterAllExceptFirst(pm.history, func(v interface{}) bool {
			_, ok := v.(*pushermsg.MetricUpdatePush)
			return ok
		})

		// Filter all file updates, keep the newest one per file
		m := lo.Uniq(lo.FilterMap(pm.history, func(v interface{}, _ int) (string, bool) {
			x, ok := v.(*pushermsg.FileStatusPush)
			if !ok {
				return "", false
			}

			return x.ID, true
		}))

		for _, hit := range m {
			pm.history = FilterAllExceptFirst(pm.history, func(v interface{}) bool {
				x, ok := v.(*pushermsg.FileStatusPush)
				return ok && x.ID == hit
			})
		}

		lo.Reverse(pm.history)

		// Increase capacity again
		pm.history = append(make([]any, 0, len(pm.history)+500), pm.history...)

		Logger.Debug().Int("before", startCount).Int("after", len(pm.history)).Msg("Websocket history compacted")

		pm.historyLock.Unlock()
	})
}

func (pm *PushManager) GetClients() []Client {
	pm.clientsLock.RLock()
	defer pm.clientsLock.RUnlock()

	// Create a static list of clients, so we can release the lock again
	clients := make([]Client, 0, len(pm.clients))
	for client := range pm.clients {
		clients = append(clients, *client)
	}

	return clients
}
