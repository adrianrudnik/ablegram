package parser

import (
	pipeline2 "github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/pusher"
	"github.com/adrianrudnik/ablegram/internal/stats"
)

type WorkerPool struct {
	workerCount         int
	inputPathsChan      <-chan *pipeline2.FilesForProcessorMsg
	outputResultChan    chan<- *pipeline2.DocumentToIndexMsg
	outputBroadcastChan chan<- interface{}
}

func NewWorkerPool(workerCount int, pathChan <-chan *pipeline2.FilesForProcessorMsg, resultChan chan<- *pipeline2.DocumentToIndexMsg, broadcastChan chan<- interface{}) *WorkerPool {
	return &WorkerPool{
		workerCount:         workerCount,
		inputPathsChan:      pathChan,
		outputResultChan:    resultChan,
		outputBroadcastChan: broadcastChan,
	}
}

func (p *WorkerPool) Run(m *stats.Metrics) {
	Logger.Info().Int("count", p.workerCount).Msg("Starting parser workers")

	for i := 0; i < p.workerCount; i++ {
		go p.doWork(m)
	}
}

func (p *WorkerPool) doWork(m *stats.Metrics) {
	for msg := range p.inputPathsChan {
		docs, err := ParseAls(msg.AbsPath, m)
		if err != nil {
			Logger.Warn().Err(err).Str("path", msg.AbsPath).Msg("Failed to parse file")

			// Notify the UI about the failure
			p.outputBroadcastChan <- pusher.NewFileStatusPush(msg.AbsPath, "failed", err.Error())

			continue
		}

		// Notify the UI about the file progress
		p.outputBroadcastChan <- pusher.NewFileStatusPush(msg.AbsPath, "processed", "")

		Logger.Debug().Str("path", msg.AbsPath).Msg("Finished processing")

		// Move the result over to the indexer pipeline
		for _, doc := range docs {
			p.outputResultChan <- doc
		}
	}
}
