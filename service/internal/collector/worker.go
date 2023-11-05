package collector

import (
	"github.com/adrianrudnik/ablegram/internal/config"
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

func (wp *WorkerPool) Run(conf *config.Config, p *stats.ProcessProgress) {
	Logger.Info().
		Int("count", wp.workerCount).
		Strs("paths", conf.Collector.SearchablePaths).
		Msg("Starting collector workers")

	// Spool up workers first
	for i := 0; i < wp.workerCount; i++ {
		go wp.doWork(conf, p)
	}

	// Pipe in paths next
	for _, path := range conf.Collector.SearchablePaths {
		wp.inPathChan <- path
	}
}

func (wp *WorkerPool) doWork(conf *config.Config, p *stats.ProcessProgress) {
	for {
		select {
		case path := <-wp.inPathChan:
			p.Add()

			err := Collect(conf, path, wp.outFilesChan, wp.pushChan)
			if err != nil {
				Logger.Warn().Err(err).Str("path", path).Msg("Failed to collect files")
			}

			p.Done()
		}
	}
}
