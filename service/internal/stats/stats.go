package stats

import (
	"github.com/adrianrudnik/ablegram/internal/config"
	"github.com/adrianrudnik/ablegram/internal/pusher"
	"github.com/samber/lo"
	"sync/atomic"
	"time"
)

type Statistics struct {
	config          *config.Config
	pushHistorySize atomic.Uint64
	counter         *Counter

	TriggerUpdate func()
}

func NewStatistics(conf *config.Config, pushChan chan<- interface{}) *Statistics {
	m := &Statistics{
		config:  conf,
		counter: NewMetric(),
	}

	// Create a debounced trigger that broadcasts the current statistics
	// towards the frontend
	m.TriggerUpdate, _ = lo.NewDebounce(50*time.Millisecond, func() {
		pushChan <- pusher.NewMetricUpdatePush(m.CollectCounters())
	})

	return m
}

func (s *Statistics) IncrementCounter(name string) {
	s.counter.Increment(name, 1)
	s.TriggerUpdate()
}

func (s *Statistics) SetCounter(name string, value uint64) {
	s.counter.Set(name, value)
	s.TriggerUpdate()
}

func (s *Statistics) CollectCounters() map[string]uint64 {
	return s.counter.Snapshot()
}
