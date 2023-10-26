package main

import (
	"flag"
	"github.com/adrianrudnik/ablegram/internal/collector"
	"github.com/adrianrudnik/ablegram/internal/config"
	"github.com/adrianrudnik/ablegram/internal/indexer"
	"github.com/adrianrudnik/ablegram/internal/parser"
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/webservice"
	"github.com/icza/gox/osx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func main() {
	// Parse flags
	noBrowserFlag := flag.Bool("no-browser", false, "Skip the automatic browser opening")
	flag.Parse()

	log.Info().Bool("no-browser", !*noBrowserFlag).Msg("Parsed executable flags")

	// Set up logging
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("App starting")

	// Let's look for a configuration within one of the folders
	config.Logger = log.With().Str("module", "config").Logger()
	appConfig := config.LoadWithDefaults("")

	// Create some channel based pipelines to pass around the different workloads
	pusherPipeline := pipeline.NewFrontendPush()
	filesPipeline := pipeline.NewFilesForProcessor()
	resultsPipeline := pipeline.NewDocumentsToIndex()

	// Metrics is responsible in keeping and communicating key metrics for the frontend
	appMetrics := stats.NewMetrics(pusherPipeline.Channel)

	// Create the indexerWorker
	indexer.Logger = log.With().Str("module", "search").Logger()
	search := indexer.NewSearch(&indexer.SearchOptions{})
	indexerWorker := indexer.NewWorker(search, resultsPipeline.Channel, pusherPipeline.Channel)
	go indexerWorker.Run(appMetrics)

	// Start the frontend push worker
	webservice.Logger = log.With().Str("module", "webservice").Logger()
	pusher := webservice.NewPushChannel(pusherPipeline.Channel)
	go pusher.Run()

	// Collector is responsible for finding files that could be parsed
	collector.Logger = log.With().Str("module", "collector").Logger()
	collectorWorkers := collector.NewWorkerPool(3, filesPipeline.Channel, pusherPipeline.Channel)
	go collectorWorkers.Run(appConfig.SearchablePaths)

	// Parser is responsible for parsing the files into results for the indexerWorker
	parser.Logger = log.With().Str("module", "parser").Logger()
	parserWorkers := parser.NewWorkerPool(5, filesPipeline.Channel, resultsPipeline.Channel, pusherPipeline.Channel)
	go parserWorkers.Run(appMetrics)

	// Try to open the default browser on the given OS
	go func() {
		if *noBrowserFlag {
			return
		}

		time.Sleep(50 * time.Millisecond)

		err := osx.OpenDefault("http://localhost:10000")
		if err != nil {
			log.Warn().Err(err).Msg("Could not open default browser")
		}
	}()

	webservice.Serve(search, pusher, ":10000")
}
