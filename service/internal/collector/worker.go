package collector

import (
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/stats"
)

type WorkerPool struct {
	workerCount  int
	paths        []string
	inPathChan   chan string
	outFilesChan chan<- *pipeline.FilesForProcessorMsg
	pushChan     chan<- interface{}
}

func NewWorkerPool(workerCount int, filesChan chan<- *pipeline.FilesForProcessorMsg, broadcastChan chan<- interface{}) *WorkerPool {
	return &WorkerPool{
		workerCount:  workerCount,
		inPathChan:   make(chan string, 100),
		outFilesChan: filesChan,
		pushChan:     broadcastChan,
	}
}

func (p *WorkerPool) Run(progress *stats.ProcessProgress, paths []string) {
	Logger.Info().
		Int("count", p.workerCount).
		Strs("paths", paths).
		Msg("Starting collector workers")

	// Spool up workers first
	for i := 0; i < p.workerCount; i++ {
		go p.doWork(progress)
	}

	// Pipe in paths next
	for _, path := range paths {
		p.inPathChan <- path
	}
}

func (p *WorkerPool) doWork(progress *stats.ProcessProgress) {
	for {
		select {
		case path := <-p.inPathChan:
			progress.Add()

			err := Collect(path, p.outFilesChan, p.pushChan)
			if err != nil {
				Logger.Warn().Err(err).Str("path", path).Msg("Failed to collect files")
			}

			progress.Done()
		}
	}
}
