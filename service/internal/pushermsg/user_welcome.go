package pushermsg

type UserWelcomePush struct {
	Type        string `json:"type"`
	ID          string `json:"id"`
	Role        string `json:"role"`
	DisplayName string `json:"display_name"`
}

func NewUserWelcomePush(id, role, displayName string) *UserWelcomePush {
	return &UserWelcomePush{
		Type:        "user_welcome",
		ID:          id,
		Role:        role,
		DisplayName: displayName,
	}
}
