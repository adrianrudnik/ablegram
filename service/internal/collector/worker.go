package collector

import (
	"github.com/adrianrudnik/ablegram/internal/pipeline"
)

type WorkerPool struct {
	workerCount         int
	paths               []string
	inputPathsChan      chan string
	outputFilesChan     chan<- *pipeline.FilesForProcessorMsg
	outputBroadcastChan chan<- interface{}
}

func NewWorkerPool(workerCount int, filesChan chan<- *pipeline.FilesForProcessorMsg, broadcastChan chan<- interface{}) *WorkerPool {
	return &WorkerPool{
		workerCount:         workerCount,
		inputPathsChan:      make(chan string, 100),
		outputFilesChan:     filesChan,
		outputBroadcastChan: broadcastChan,
	}
}

func (p *WorkerPool) Run(paths []string) {
	Logger.Info().
		Int("count", p.workerCount).
		Strs("paths", paths).
		Msg("Starting collector workers")

	// Spool up workers first
	for i := 0; i < p.workerCount; i++ {
		go p.doWork()
	}

	// Pipe in paths next
	for _, path := range paths {
		p.inputPathsChan <- path
	}
}

func (p *WorkerPool) doWork() {
	for {
		select {
		case path := <-p.inputPathsChan:
			err := Collect(path, p.outputFilesChan, p.outputBroadcastChan)
			if err != nil {
				Logger.Warn().Err(err).Str("path", path).Msg("Failed to collect files")
			}
		}
	}
}
