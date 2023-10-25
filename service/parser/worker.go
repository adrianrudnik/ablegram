package parser

import (
	"github.com/adrianrudnik/ablegram/pipeline"
	"github.com/adrianrudnik/ablegram/pusher"
)

type WorkerPool struct {
	workerCount         int
	inputPathsChan      <-chan *pipeline.FilesForProcessorMsg
	outputResultChan    chan<- *pipeline.ResultToIndexMsg
	outputBroadcastChan chan<- interface{}
}

func NewWorkerPool(workerCount int, pathChan <-chan *pipeline.FilesForProcessorMsg, resultChan chan<- *pipeline.ResultToIndexMsg, broadcastChan chan<- interface{}) *WorkerPool {
	return &WorkerPool{
		workerCount:         workerCount,
		inputPathsChan:      pathChan,
		outputResultChan:    resultChan,
		outputBroadcastChan: broadcastChan,
	}
}

func (p *WorkerPool) Run() {
	Logger.Info().Int("count", p.workerCount).Msg("Starting parser workers")

	for i := 0; i < p.workerCount; i++ {
		go p.doWork()
	}
}

func (p *WorkerPool) doWork() {
	for msg := range p.inputPathsChan {
		_, err := ParseAls(msg.AbsPath)
		if err != nil {
			Logger.Warn().Err(err).Str("path", msg.AbsPath).Msg("Failed to parse file")

			// Notify the UI about the failure
			p.outputBroadcastChan <- pusher.NewFileStatusPush(msg.AbsPath, "failed", err.Error())

			continue
		}

		// Notify the UI about the file progress
		p.outputBroadcastChan <- pusher.NewFileStatusPush(msg.AbsPath, "processed", "")

		Logger.Info().Str("path", msg.AbsPath).Msg("Finished processing")

		// Move the result over to the indexer pipeline
		p.outputResultChan <- &pipeline.ResultToIndexMsg{}
	}
}
