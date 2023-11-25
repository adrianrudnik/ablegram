package pushermsg

// AboutYouPush is a message sent to a single client to inform him about his own ID in the backend.
// This is later used to upgrade websockets on a guest => admin migration to upgrade the pusher client as well.
type AboutYouPush struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// GetUsers ensures this message is only routed towards the given client ID.
func (p *AboutYouPush) GetUsers() map[string]bool {
	return map[string]bool{p.ID: true}
}

// KeepInHistory ensures the message is not kept in history.
func (p *AboutYouPush) KeepInHistory() bool {
	return false
}

func NewAboutYouPush(id string) *AboutYouPush {
	return &AboutYouPush{
		Type: "about_you",
		ID:   id,
	}
}
