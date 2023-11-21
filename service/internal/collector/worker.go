package collector

import (
	"github.com/adrianrudnik/ablegram/internal/config"
	"github.com/adrianrudnik/ablegram/internal/parser"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/tagger"
	"github.com/adrianrudnik/ablegram/internal/workload"
)

type WorkerPool struct {
	config   *config.Config
	stat     *stats.Statistics
	progress *stats.ProcessProgress
	tc       *tagger.TagCollector

	indexChan chan<- *workload.DocumentPayload
	pushChan  chan<- interface{}
}

func NewWorkerPool(
	conf *config.Config,
	stat *stats.Statistics,
	progress *stats.ProcessProgress,
	tc *tagger.TagCollector,

	indexChan chan<- *workload.DocumentPayload,
	pushChan chan<- interface{},
) *WorkerPool {
	return &WorkerPool{
		config:   conf,
		stat:     stat,
		progress: progress,
		tc:       tc,

		indexChan: indexChan,
		pushChan:  pushChan,
	}
}

// Run starts the worker pool, which will spawn workers for each target
func (wp *WorkerPool) Run() {
	Logger.Info().
		Int("targets", len(wp.config.Collector.Targets)).
		Msg("Starting collector working pool")

	for _, t := range wp.config.Collector.Targets {
		go wp.spawnTargetWorker(wp.config, t)
	}
}

func (wp *WorkerPool) spawnTargetWorker(
	conf *config.Config,
	target config.CollectorTarget,
) {
	// Every target worker has its own file channel to collect files from
	pathChan := make(chan *workload.FilePayload, 1000)

	// Every target worker has its own parser worker pool
	pwp := parser.NewWorkerPool(
		conf,
		wp.stat,
		wp.progress,
		wp.tc,
		&target,

		pathChan,     // to receive files to parse from
		wp.indexChan, // to send parsed results to the indexer
		wp.pushChan,  // to notify the frontend about the progress
	)

	pwp.Run()

	// Start the collection process for this target
	wp.progress.Add()
	err := Collect(conf, &target, pathChan, wp.pushChan)
	wp.progress.Done()
	if err != nil {
		Logger.Warn().Err(err).Str("path", target.Uri).Msg("Failed to collect files")
	}
}
