package indexer

import (
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/workload"
	"github.com/rs/zerolog/log"
)

type Worker struct {
	search   *Search
	docChan  <-chan *workload.DocumentPayload
	pushChan chan<- interface{}
}

func NewWorker(search *Search, docChan <-chan *workload.DocumentPayload, broadcastChan chan<- interface{}) *Worker {
	return &Worker{
		search:   search,
		docChan:  docChan,
		pushChan: broadcastChan,
	}
}

func (p *Worker) Run(progress *stats.ProcessProgress, m *stats.Statistics) {
	Logger.Info().Msg("Starting index batch worker")

	go p.doWork(progress, m)
}

func (p *Worker) doWork(progress *stats.ProcessProgress, stat *stats.Statistics) {
	for {
		select {
		case msg := <-p.docChan:
			progress.Add()

			err := p.search.Index.Index(msg.Id, msg.Document)

			progress.Done()
			if err != nil {
				log.Error().Err(err).Str("document", msg.Id).Msg("Failed to add document to batch")

				continue
			}

			docCount, err := p.search.Index.DocCount()
			if err != nil {
				log.Warn().Err(err).Msg("Failed to retrieve document count from index")
				continue
			}

			stat.SetCounter(IndexDocument, docCount)

			continue
		}
	}
}
