package suggest

import "github.com/google/uuid"

type Entry struct {
	ID     uuid.UUID `json:"id"`
	Owner  uuid.UUID `json:"owner"`
	Target string    `json:"target"`
}
