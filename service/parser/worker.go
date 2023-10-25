package parser

import "github.com/adrianrudnik/ablegram/pipeline"

type WorkerPool struct {
	workerCount      int
	inputPathsChan   <-chan *pipeline.FilesForProcessorMsg
	outputResultChan chan<- *pipeline.ResultToIndexMsg
}

func NewWorkerPool(workerCount int, pathChan <-chan *pipeline.FilesForProcessorMsg, resultChan chan<- *pipeline.ResultToIndexMsg) *WorkerPool {
	return &WorkerPool{
		workerCount:      workerCount,
		inputPathsChan:   pathChan,
		outputResultChan: resultChan,
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
		}

		p.outputResultChan <- &pipeline.ResultToIndexMsg{}
	}
}
