package webservice

import (
	"github.com/adrianrudnik/ablegram/internal/pusher"
)

func canClientReceive(msg any, client *PushClient) bool {
	// Guard role routed messages
	if msg, ok := msg.(pusher.RoleRoutedMessage); ok {
		if msg.GetRole() != client.Role {
			return false
		}
	}

	// Guard user routed messages
	if msg, ok := msg.(pusher.UserRouterMessage); ok {
		if _, ok := msg.GetUsers()[client.ID]; !ok {
			return false
		}
	}

	return true
}
