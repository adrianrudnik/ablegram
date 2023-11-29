package suggest

import (
	"github.com/google/uuid"
	"net/url"
	"sync"
)

type Endpoint struct {
	ID    uuid.UUID         `json:"id"`
	Label string            `json:"label"`
	Uri   string            `json:"uri"`
	Extra map[string]string `json:"extra"`
	mutex sync.RWMutex
}

func (e *Endpoint) ParsedUri() (*url.URL, error) {
	return url.Parse(e.Uri)
}

func (e *Endpoint) SetExtra(key, value string) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	e.Extra[key] = value
}

func (e *Endpoint) RemoveExtra(key string) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	delete(e.Extra, key)
}

func (e *Endpoint) GetExtra(key string) string {
	e.mutex.RLock()
	defer e.mutex.RUnlock()

	return e.Extra[key]
}
