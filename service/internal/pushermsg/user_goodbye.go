package pushermsg

import "github.com/google/uuid"

type UserGoodbyePush struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// KeepInHistory ensures the message is not kept in history, as there is no need to, new or reconnecting clients are
// only interested in the current state.
func (p *UserGoodbyePush) KeepInHistory() bool {
	return false
}

func NewUserGoodbyePush(id uuid.UUID) *UserGoodbyePush {
	return &UserGoodbyePush{
		Type: "user_goodbye",
		ID:   id.String(),
	}
}
