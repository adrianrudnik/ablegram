package pusher

import "github.com/adrianrudnik/ablegram/internal/access"

func (pm *PushManager) canClientReceive(msg any, client *Client) bool {
	// Guard role routed messages
	if msg, ok := msg.(RoleRoutedMessage); ok {
		user := pm.users.Get(client.UserID)
		if msg.GetRole() != user.Role {
			return false
		}
	}

	// Guard user routed messages
	if msg, ok := msg.(UserRouterMessage); ok {
		user := pm.users.Get(client.UserID)
		if _, ok := msg.GetUsers()[user.UserID]; !ok {
			return false
		}
	}

	return true
}

func (pm *PushManager) reduceMessageForClient(msg any, client *Client) any {
	if v, ok := msg.(FilteredMessage); ok && pm.users.Get(client.UserID).Role == access.GuestRole {
		return v.FilteredVariant()
	} else {
		return msg
	}
}
