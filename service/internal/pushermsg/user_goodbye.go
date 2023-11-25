package pushermsg

type UserGoodbyePush struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

func NewUserGoodbyePush(id string) *UserGoodbyePush {
	return &UserGoodbyePush{
		Type: "user_goodbye",
		ID:   id,
	}
}
