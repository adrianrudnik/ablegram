package pushermsg

type UserWelcomePush struct {
	ID          string `json:"id"`
	Role        string `json:"type"`
	DisplayName string `json:"display_name"`
}

func NewUserWelcomePush(id, role, displayName string) *UserWelcomePush {
	return &UserWelcomePush{
		ID:          id,
		Role:        role,
		DisplayName: displayName,
	}
}
