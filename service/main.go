package main

import (
	"flag"
	collector2 "github.com/adrianrudnik/ablegram/internal/collector"
	"github.com/adrianrudnik/ablegram/internal/config"
	parser2 "github.com/adrianrudnik/ablegram/internal/parser"
	pipeline2 "github.com/adrianrudnik/ablegram/internal/pipeline"
	search2 "github.com/adrianrudnik/ablegram/internal/search"
	"github.com/adrianrudnik/ablegram/internal/stats"
	webservice2 "github.com/adrianrudnik/ablegram/internal/webservice"
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
	pusherPipeline := pipeline2.NewFrontendPush()
	filesPipeline := pipeline2.NewFilesForProcessor()
	resultsPipeline := pipeline2.NewDocumentsToIndex()

	// Metrics is responsible in keeping and communicating key metrics for the frontend
	appMetrics := stats.NewMetrics(pusherPipeline.Channel)

	// Create the indexer
	search2.Logger = log.With().Str("module", "search").Logger()
	appSearch := search2.NewSearch(&search2.SearchOptions{})
	indexer := search2.NewWorker(appSearch, resultsPipeline.Channel, pusherPipeline.Channel)
	go indexer.Run(appMetrics)

	// Start the frontend push worker
	webservice2.Logger = log.With().Str("module", "webservice").Logger()
	pusher := webservice2.NewPushChannel(pusherPipeline.Channel)
	go pusher.Run()

	// Collector is responsible for finding files that could be parsed
	collector2.Logger = log.With().Str("module", "collector").Logger()
	collectorWorkers := collector2.NewWorkerPool(3, filesPipeline.Channel, pusherPipeline.Channel)
	go collectorWorkers.Run(appConfig.SearchablePaths)

	// Parser is responsible for parsing the files into results for the indexer
	parser2.Logger = log.With().Str("module", "parser").Logger()
	parserWorkers := parser2.NewWorkerPool(5, filesPipeline.Channel, resultsPipeline.Channel, pusherPipeline.Channel)
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

	webservice2.Serve(pusher, ":10000")
}
