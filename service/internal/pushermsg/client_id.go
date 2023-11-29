package pushermsg

import "github.com/google/uuid"

// ClientIdPush is a message that announces the current client ID over the connected websocket the ID belongs to.
type ClientIdPush struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// GetUsers ensures this message is only routed towards the given client ID.
func (p *ClientIdPush) GetUsers() map[string]bool {
	return map[string]bool{p.ID: true}
}

// KeepInHistory ensures the message is not kept in history.
func (p *ClientIdPush) KeepInHistory() bool {
	return false
}

func NewClientIdPush(id uuid.UUID) *ClientIdPush {
	return &ClientIdPush{
		Type: "client_id",
		ID:   id.String(),
	}
}
