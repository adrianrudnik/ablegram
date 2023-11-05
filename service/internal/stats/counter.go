package stats

import "sync"

const (
	FileValid   = "FileValid" // c0 == 0
	FileInvalid = "FileInvalid"
)

type Counter struct {
	mu       sync.RWMutex
	counters map[string]uint64
}

func NewMetric() *Counter {
	return &Counter{
		counters: make(map[string]uint64),
	}
}

func (m *Counter) Increment(k string, v uint64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.counters[k]; !ok {
		m.counters[k] = 0
	}

	m.counters[k] += v
}

func (m *Counter) Set(k string, v uint64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.counters[k]; !ok {
		m.counters[k] = 0
	}

	m.counters[k] = v
}

func (m *Counter) Snapshot() map[string]uint64 {
	m.mu.RLock()
	defer m.mu.RUnlock()

	snapshot := make(map[string]uint64)
	for k, v := range m.counters {
		snapshot[k] = v
	}

	return snapshot
}
