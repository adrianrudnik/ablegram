package pushermsg

type UserWelcomePush struct {
	Type        string `json:"type"`
	ID          string `json:"id"`
	Role        string `json:"role"`
	DisplayName string `json:"display_name"`
	IP          string `json:"ip"`
}

func (p *UserWelcomePush) FilteredVariant() interface{} {
	return &UserWelcomePush{
		Type:        p.Type,
		ID:          p.ID,
		Role:        p.Role,
		DisplayName: p.DisplayName,
		IP:          "",
	}
}

// KeepInHistory ensures the message is not kept in history, as we can send the idempotent list on connect
func (p *UserWelcomePush) KeepInHistory() bool {
	return false
}

func NewUserWelcomePush(id, role, displayName, ip string) *UserWelcomePush {
	return &UserWelcomePush{
		Type:        "user_welcome",
		ID:          id,
		Role:        role,
		DisplayName: displayName,
		IP:          ip,
	}
}
