package parser

import (
	"github.com/adrianrudnik/ablegram/internal/config"
	pipeline2 "github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/pusher"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"time"
)

type WorkerPool struct {
	config           *config.ParserConfig
	workerCount      int
	inputPathsChan   <-chan *pipeline2.FilesForProcessorMsg
	outputResultChan chan<- *pipeline2.DocumentToIndexMsg
	pushChan         chan<- interface{}
}

func NewWorkerPool(c *config.ParserConfig, pathChan <-chan *pipeline2.FilesForProcessorMsg, resultChan chan<- *pipeline2.DocumentToIndexMsg, pushChan chan<- interface{}) *WorkerPool {
	return &WorkerPool{
		config:           c,
		inputPathsChan:   pathChan,
		outputResultChan: resultChan,
		pushChan:         pushChan,
	}
}

func (p *WorkerPool) Run(progress *stats.ProcessProgress, m *stats.Metrics) {
	Logger.Info().Int("count", p.workerCount).Msg("Starting parser workers")

	for i := 0; i < p.config.WorkerCount; i++ {
		go p.doWork(progress, m)
	}
}

func (p *WorkerPool) doWork(progress *stats.ProcessProgress, m *stats.Metrics) {
	for msg := range p.inputPathsChan {
		// Add possible delay, for debugging or to simulate a slower system
		if p.config.WorkerDelayInMs > 0 {
			time.Sleep(time.Duration(p.config.WorkerDelayInMs) * time.Millisecond)
		}

		progress.Add()
		docs, err := ParseAls(msg.AbsPath, m)

		progress.Done()
		if err != nil {
			Logger.Warn().Err(err).Str("path", msg.AbsPath).Msg("Failed to parse file")

			// Notify the UI about the failure
			p.pushChan <- pusher.NewFileStatusPush(msg.AbsPath, "failed", err.Error())

			continue
		}

		// Notify the UI about the file progress
		p.pushChan <- pusher.NewFileStatusPush(msg.AbsPath, "processed", "")

		Logger.Debug().Str("path", msg.AbsPath).Msg("Finished processing")

		// Move the result over to the indexer pipeline
		for _, doc := range docs {
			p.outputResultChan <- doc
		}
	}
}
