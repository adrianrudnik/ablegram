package pusher

func CanClientReceive(msg any, client *PushClient) bool {
	// Guard role routed messages
	if msg, ok := msg.(RoleRoutedMessage); ok {
		if msg.GetRole() != client.Role {
			return false
		}
	}

	// Guard user routed messages
	if msg, ok := msg.(UserRouterMessage); ok {
		if _, ok := msg.GetUsers()[client.ID]; !ok {
			return false
		}
	}

	return true
}
