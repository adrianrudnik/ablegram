package pushermsg

import (
	"github.com/google/uuid"
	"net"
)

type UserWelcomePush struct {
	Type        string `json:"type"`
	ID          string `json:"id"`
	Role        string `json:"role"`
	DisplayName string `json:"display_name"`
	IP          string `json:"ip"`
}

func (p *UserWelcomePush) FilteredVariant() interface{} {
	v := *p
	v.IP = ""

	return v
}

// KeepInHistory ensures the message is not kept in history, as we can send the idempotent list on connect
func (p *UserWelcomePush) KeepInHistory() bool {
	return false
}

func NewUserWelcomePush(clientIp net.IP, userId uuid.UUID, userRole string, userDisplayName string) *UserWelcomePush {
	return &UserWelcomePush{
		Type:        "user_welcome",
		ID:          userId.String(),
		Role:        userRole,
		DisplayName: userDisplayName,
		IP:          clientIp.String(),
	}
}
