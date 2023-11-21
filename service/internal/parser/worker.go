package parser

import (
	"github.com/adrianrudnik/ablegram/internal/config"
	"github.com/adrianrudnik/ablegram/internal/pusher"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
	"github.com/adrianrudnik/ablegram/internal/workload"
	"time"
)

type WorkerPool struct {
	config   *config.Config
	stat     *stats.Statistics
	progress *stats.ProcessProgress
	tc       *tagger.TagCollector
	target   *config.CollectorTarget

	pathChan  <-chan *workload.FilePayload
	indexChan chan<- *workload.DocumentPayload
	pushChan  chan<- interface{}
}

func NewWorkerPool(
	conf *config.Config,
	stat *stats.Statistics,
	progress *stats.ProcessProgress,
	tc *tagger.TagCollector,
	target *config.CollectorTarget,

	pathChan <-chan *workload.FilePayload,
	indexChan chan<- *workload.DocumentPayload,
	pushChan chan<- interface{},
) *WorkerPool {
	return &WorkerPool{
		config:   conf,
		stat:     stat,
		progress: progress,
		tc:       tc,
		target:   target,

		pathChan:  pathChan,
		indexChan: indexChan,
		pushChan:  pushChan,
	}
}

func (p *WorkerPool) Run() {
	Logger.Info().
		Str("performance-mode", p.target.ParserPerformance).
		Str("collector-id", p.target.ID).
		Str("collector-uri", p.target.Uri).
		Msg("Starting parser workers for collector target")

	// Decide on a worker count based on the requested performance mode
	var c int

	switch p.target.ParserPerformance {
	case "low":
		c = 1
	case "high":
		c = 5
	default:
		c = 3
	}

	for i := 0; i < c; i++ {
		go p.doWork()
	}
}

func (p *WorkerPool) doWork() {
	for msg := range p.pathChan {
		// Add possible delay, for debugging or to simulate a slower system
		if p.target.ParserWorkerDelay > 0 {
			time.Sleep(time.Duration(p.target.ParserWorkerDelay) * time.Millisecond)
		}

		p.progress.Add()
		docs, err := ParseAls(p.stat, p.tc, msg.AbsPath)
		p.progress.Done()

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
			// skip empty ones, parsers are allowed to return nil documents if they deem them invalid
			if doc == nil {
				continue
			}

			p.indexChan <- doc
		}
	}
}
