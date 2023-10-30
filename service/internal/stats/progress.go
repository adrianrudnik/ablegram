package stats

import (
	"github.com/adrianrudnik/ablegram/internal/pusher"
	"github.com/samber/lo"
	"sync/atomic"
	"time"
)

// ProcessProgress is a helper to track the progress of all processors
// involved with parsing and indexing possible files.
// It emits a broadcast message to the UI whenever the progress changes,
// in a debounced way, so the UI can decide if the service is working (n > 0)
// or all tasks have finished (n == 0).
type ProcessProgress struct {
	progressCount atomic.Int64       // Expose the counter to UI
	pushChan      chan<- interface{} // Broadcast updates
	pushTrigger   func()             // Triggers a debounced push broadcast
}

func NewProcessProgress(pushChan chan<- interface{}) *ProcessProgress {
	p := &ProcessProgress{
		pushChan: pushChan,
	}

	p.pushTrigger, _ = lo.NewDebounce(50*time.Millisecond, func() {
		p.pushChan <- pusher.NewProcessingStatusPush(p.progressCount.Load())
		Logger.Debug().Int64("routines", p.progressCount.Load()).Msg("Processing progress updated")

		if p.progressCount.Load() == 0 {
			Logger.Info().Msg("Processing finished")
		}
	})

	return p
}

func (p *ProcessProgress) Add() {
	p.progressCount.Add(1)
	p.pushTrigger()
}

func (p *ProcessProgress) Done() {
	p.progressCount.Add(-1)
	p.pushTrigger()
}

func (p *ProcessProgress) IsInProgress() bool {
	return p.progressCount.Load() > 0
}
