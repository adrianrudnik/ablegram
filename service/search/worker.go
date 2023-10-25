package search

import (
	"github.com/adrianrudnik/ablegram/pipeline"
	"github.com/adrianrudnik/ablegram/pusher"
	"github.com/rs/zerolog/log"
)

type Worker struct {
	search              *Search
	inputDocChan        <-chan *pipeline.DocumentToIndexMsg
	outputBroadcastChan chan<- interface{}
}

func NewWorker(search *Search, docChan <-chan *pipeline.DocumentToIndexMsg, broadcastChan chan<- interface{}) *Worker {
	return &Worker{
		search:              search,
		inputDocChan:        docChan,
		outputBroadcastChan: broadcastChan,
	}
}

func (p *Worker) Run() {
	Logger.Info().Msg("Starting index batch worker")

	go p.doWork()
}

func (p *Worker) doWork() {
	for {
		select {
		case msg := <-p.inputDocChan:
			err := p.search.Index.Index(msg.Id, msg.Document)
			if err != nil {
				log.Error().Err(err).Str("document", msg.Id).Msg("Failed to add document to batch")
			}

			log.Debug().Str("document", msg.Id).Msg("Document indexed")

			docCount, err := p.search.Index.DocCount()
			if err != nil {
				log.Warn().Err(err).Msg("Failed to retrieve document count from index")
				continue
			}

			p.outputBroadcastChan <- pusher.NewIndexStatusPush(docCount)

			continue
		}
	}
}
