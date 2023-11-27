package pusher

import (
	"github.com/google/uuid"
	"net"
)

// UserClient represents a user that is currently connected to the system.
// A client can have multiple connections (e.g. multiple tabs).
type UserClient struct {
	ClientId uuid.UUID `json:"id"`
	ClientIP net.IP    `json:"ip"`
	UserId   uuid.UUID `json:"user_id"`
}
