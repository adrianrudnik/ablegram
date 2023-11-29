package suggest

import (
	"github.com/google/uuid"
	"sync"
)

type EndpointList struct {
	mutex     sync.RWMutex
	endpoints map[uuid.UUID]*Endpoint
}

func NewList() *EndpointList {
	return &EndpointList{
		endpoints: make(map[uuid.UUID]*Endpoint),
	}
}

func (s *EndpointList) All() []Endpoint {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var suggestions []Endpoint

	for _, endpoint := range s.endpoints {
		suggestions = append(suggestions, *endpoint)
	}

	return suggestions
}

func (s *EndpointList) Add(endpoint *Endpoint) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.endpoints[endpoint.ID] = endpoint
}

func (s *EndpointList) Get(id uuid.UUID) *Endpoint {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if endpoint, ok := s.endpoints[id]; ok {
		return endpoint
	}

	return nil
}

func (s *EndpointList) Delete(id uuid.UUID) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.endpoints, id)
}
