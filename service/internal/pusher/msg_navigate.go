package pusher

type ForceNavigatePush struct {
	Type   string `json:"type"`
	Target string `json:"target"`
}

func (p *ForceNavigatePush) KeepInHistory() bool {
	return false
}

func NewForceNavigatePush(url string) *ForceNavigatePush {
	return &ForceNavigatePush{
		Type:   "force_navigate",
		Target: url,
	}
}
