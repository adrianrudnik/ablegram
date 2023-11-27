package pushermsg

import "github.com/google/uuid"

type ClientGoodbyePush struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// KeepInHistory ensures the message is not kept in history, as there is no need to, new or reconnecting clients are
// only interested in the current state.
func (p *ClientGoodbyePush) KeepInHistory() bool {
	return false
}

func NewClientGoodbyePush(id uuid.UUID) *ClientGoodbyePush {
	return &ClientGoodbyePush{
		Type: "client_goodbye",
		ID:   id.String(),
	}
}
