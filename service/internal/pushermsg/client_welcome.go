package pushermsg

import (
	"github.com/google/uuid"
	"net"
)

// ClientWelcomePush is a message that communicates a new user client connecting to all other connected users.
type ClientWelcomePush struct {
	Type        string `json:"type"`
	ID          string `json:"id"`
	Role        string `json:"role"`
	DisplayName string `json:"display_name"`
	IP          string `json:"ip"`
}

func (p *ClientWelcomePush) FilteredVariant() interface{} {
	v := *p
	v.IP = ""

	return v
}

// KeepInHistory ensures the message is not kept in history, as we can send the idempotent list on connect
func (p *ClientWelcomePush) KeepInHistory() bool {
	return false
}

func NewClientWelcomePush(clientId uuid.UUID, clientIp net.IP, userId uuid.UUID, userRole string, userDisplayName string) *UserClientPush {
	return &UserClientPush{
		Type: "client_welcome",

		ClientId: clientId.String(),
		ClientIP: clientIp.String(),

		UserId:          userId.String(),
		UserRole:        userRole,
		UserDisplayName: userDisplayName,
	}
}
