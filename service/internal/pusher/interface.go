package pusher

import "github.com/google/uuid"

type RoleRoutedMessage interface {
	GetRole() string
}

type UserRouterMessage interface {
	GetUsers() map[uuid.UUID]bool
}

type RecordMessage interface {
	KeepInHistory() bool
}

type FilteredMessage interface {
	FilteredVariant() interface{}
}
