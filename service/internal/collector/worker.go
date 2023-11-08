package collector

import (
	"github.com/adrianrudnik/ablegram/internal/config"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/workload"
	"time"
)

type WorkerPool struct {
	config       *config.Config
	paths        []string
	inPathChan   chan string
	outFilesChan chan<- *workload.FilePayload
	pushChan     chan<- interface{}
}

func NewWorkerPool(
	conf *config.Config,
	filesChan chan<- *workload.FilePayload,
	broadcastChan chan<- interface{},
) *WorkerPool {
	return &WorkerPool{
		config:       conf,
		inPathChan:   make(chan string, 100),
		outFilesChan: filesChan,
		pushChan:     broadcastChan,
	}
}

func (wp *WorkerPool) Run(p *stats.ProcessProgress) {
	Logger.Info().
		Int("count", wp.config.Collector.WorkerCount).
		Strs("paths", wp.config.Collector.SearchablePaths).
		Msg("Starting collector workers")

	// Spool up workers first
	for i := 0; i < wp.config.Collector.WorkerCount; i++ {
		go wp.doWork(wp.config, p)
	}

	// Pipe in paths next
	for _, path := range wp.config.Collector.SearchablePaths {
		wp.inPathChan <- path
	}
}

func (wp *WorkerPool) doWork(conf *config.Config, progress *stats.ProcessProgress) {
	for path := range wp.inPathChan {
		// Add possible delay, for debugging or to simulate a slower system
		if wp.config.Collector.WorkerDelayInMs > 0 {
			time.Sleep(time.Duration(wp.config.Collector.WorkerDelayInMs) * time.Millisecond)
		}

		progress.Add()

		err := Collect(conf, path, wp.outFilesChan, wp.pushChan)
		if err != nil {
			Logger.Warn().Err(err).Str("path", path).Msg("Failed to collect files")
		}

		progress.Done()
	}
}
