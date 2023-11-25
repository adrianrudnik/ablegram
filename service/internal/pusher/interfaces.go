package pusher

type RoleRoutedMessage interface {
	GetRole() string
}

type UserRouterMessage interface {
	GetUsers() map[string]bool
}

type RecordMessage interface {
	KeepInHistory() bool
}
