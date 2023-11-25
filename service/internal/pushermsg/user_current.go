package pushermsg

// UserCurrentPush is a message that communicates a user's current state to a single client.
type UserCurrentPush struct {
	Type        string `json:"type"`
	To          string `json:"to"`
	ID          string `json:"id"`
	Role        string `json:"role"`
	DisplayName string `json:"display_name"`
	IP          string `json:"ip"`
}

// GetUsers ensures this message is only routed towards the given client ID.
func (p *UserCurrentPush) GetUsers() map[string]bool {
	return map[string]bool{p.To: true}
}

// KeepInHistory ensures the message is not kept in history.
func (p *UserCurrentPush) KeepInHistory() bool {
	return false
}

func (p *UserCurrentPush) FilteredVariant() interface{} {
	return &UserCurrentPush{
		Type:        p.Type,
		To:          p.To,
		ID:          p.ID,
		Role:        p.Role,
		DisplayName: p.DisplayName,
		IP:          "",
	}
}

func NewUserCurrentPush(to, id, role, displayName, ip string) *UserCurrentPush {
	return &UserCurrentPush{
		Type:        "user_current",
		To:          to,
		ID:          id,
		Role:        role,
		DisplayName: displayName,
		IP:          ip,
	}
}
