package suggest

import (
	"github.com/google/uuid"
	"sync"
)

type List struct {
	mutex       sync.RWMutex
	suggestions map[uuid.UUID]Entry
}

func NewList() *List {
	return &List{
		suggestions: make(map[uuid.UUID]Entry),
	}
}

func (s *List) All() []Entry {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var suggestions []Entry

	for _, suggestion := range s.suggestions {
		suggestions = append(suggestions, suggestion)
	}

	return suggestions
}

func (s *List) Add(suggestion *Entry) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.suggestions[suggestion.ID] = *suggestion
}

func (s *List) Get(id uuid.UUID) *Entry {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if suggestion, ok := s.suggestions[id]; ok {
		return &suggestion
	}

	return nil
}

func (s *List) Delete(id uuid.UUID) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.suggestions, id)
}
