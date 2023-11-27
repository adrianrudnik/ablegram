package pushermsg

type SuggestionCreated struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

func (p *SuggestionCreated) GetRole() string {
	return "admin"
}

func NewSuggestionCreated(to, id, role, displayName, ip string) *SuggestionCreated {
	return &SuggestionCreated{
		Type: "suggestion_created",
		ID:   id,
	}
}
