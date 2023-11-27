package pushermsg

import (
	"github.com/google/uuid"
	"net"
)

// UserClientPush communicates a specific state of a client to a specifc user.
// Used to send the initial client connection state towards new websocket connections.
// Is separate to the ClientWelcomePush as it is only sent to a specific user and is guarded by the GetUsers() interface.
type UserClientPush struct {
	Type string `json:"type"`

	ClientId string `json:"id"`
	ClientIP string `json:"ip"`

	UserId          string `json:"user_id"`
	UserRole        string `json:"user_role"`
	UserDisplayName string `json:"user_display_name"`

	toUserId string
}

// GetUsers ensures this message is only routed towards the given client ID.
func (p *UserClientPush) GetUsers() map[string]bool {
	return map[string]bool{p.toUserId: true}
}

// KeepInHistory ensures the message is not kept in history.
func (p *UserClientPush) KeepInHistory() bool {
	return false
}

func (p *UserClientPush) FilteredVariant() interface{} {
	v := *p
	v.ClientIP = ""

	return v
}

func NewUserClientPush(toUserId uuid.UUID, clientId uuid.UUID, clientIp net.IP, userId uuid.UUID, userRole string, userDisplayName string) *UserClientPush {
	return &UserClientPush{
		Type:     "user_client",
		toUserId: toUserId.String(),

		ClientId: clientId.String(),
		ClientIP: clientIp.String(),

		UserId:          userId.String(),
		UserRole:        userRole,
		UserDisplayName: userDisplayName,
	}
}
