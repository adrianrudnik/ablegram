package parser

import (
	"github.com/adrianrudnik/ablegram/internal/config"
	pipeline2 "github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/pusher"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"time"
)

type WorkerPool struct {
	config           *config.Config
	workerCount      int
	inputPathsChan   <-chan *pipeline2.FilesForProcessorMsg
	outputResultChan chan<- *pipeline2.DocumentToIndexMsg
	pushChan         chan<- interface{}
}

func NewWorkerPool(conf *config.Config, pathChan <-chan *pipeline2.FilesForProcessorMsg, resultChan chan<- *pipeline2.DocumentToIndexMsg, pushChan chan<- interface{}) *WorkerPool {
	return &WorkerPool{
		config:           conf,
		inputPathsChan:   pathChan,
		outputResultChan: resultChan,
		pushChan:         pushChan,
	}
}

func (p *WorkerPool) Run(progress *stats.ProcessProgress, stat *stats.Statistics) {
	Logger.Info().Int("count", p.workerCount).Msg("Starting parser workers")

	for i := 0; i < p.config.ParserConfig.WorkerCount; i++ {
		go p.doWork(progress, stat)
	}
}

func (p *WorkerPool) doWork(progress *stats.ProcessProgress, m *stats.Statistics) {
	for msg := range p.inputPathsChan {
		// Add possible delay, for debugging or to simulate a slower system
		if p.config.ParserConfig.WorkerDelayInMs > 0 {
			time.Sleep(time.Duration(p.config.ParserConfig.WorkerDelayInMs) * time.Millisecond)
		}

		progress.Add()
		docs, err := ParseAls(m, msg.AbsPath)

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
