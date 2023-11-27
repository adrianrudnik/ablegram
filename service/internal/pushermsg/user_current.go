package pushermsg

import (
	"github.com/google/uuid"
	"net"
)

// UserCurrentPush is a message that communicates a user's current state to a single client.
type UserCurrentPush struct {
	Type string `json:"type"`

	ClientId string `json:"client_id"`
	ClientIP string `json:"client_ip"`

	UserId          string `json:"user_id"`
	UserRole        string `json:"user_role"`
	UserDisplayName string `json:"user_display_name"`

	toUserId string
}

// GetUsers ensures this message is only routed towards the given client ID.
func (p *UserCurrentPush) GetUsers() map[string]bool {
	return map[string]bool{p.toUserId: true}
}

// KeepInHistory ensures the message is not kept in history.
func (p *UserCurrentPush) KeepInHistory() bool {
	return false
}

func (p *UserCurrentPush) FilteredVariant() interface{} {
	v := *p
	v.ClientIP = ""

	return v
}

func NewUserCurrentPush(toUserId uuid.UUID, clientId uuid.UUID, clientIp net.IP, userId uuid.UUID, userRole string, userDisplayName string) *UserCurrentPush {
	return &UserCurrentPush{
		Type:     "user_current",
		toUserId: toUserId.String(),

		ClientId: clientId.String(),
		ClientIP: clientIp.String(),

		UserId:          userId.String(),
		UserRole:        userRole,
		UserDisplayName: userDisplayName,
	}
}
