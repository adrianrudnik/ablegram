package main

import (
	"github.com/adrianrudnik/ablegram/collector"
	"github.com/adrianrudnik/ablegram/config"
	"github.com/adrianrudnik/ablegram/parser"
	"github.com/adrianrudnik/ablegram/pipeline"
	"github.com/adrianrudnik/ablegram/webservice"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	// Set up logging
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Debug().Msg("App starting")

	// Let's look for a configuration within one of the folders
	config.Logger = log.With().Str("module", "config").Logger()
	appConfig := config.LoadWithDefaults("")

	// Create some channel based pipelines to pass around the different workloads
	pusherPipeline := pipeline.NewFrontendPush()
	filesPipeline := pipeline.NewFilesForProcessor()
	resultsPipeline := pipeline.NewResultsToIndex()

	// Start the frontend push worker
	webservice.Logger = log.With().Str("module", "webservice").Logger()
	pusher := webservice.NewPushChannel(pusherPipeline.Channel)
	go pusher.Run()

	// Collector is responsible for finding files that could be parsed
	collector.Logger = log.With().Str("module", "collector").Logger()
	collectorWorkers := collector.NewWorkerPool(3, filesPipeline.Channel, pusherPipeline.Channel)
	go collectorWorkers.Run(appConfig.SearchablePaths)

	// Parser is responsible for parsing the files into results for the indexer
	parser.Logger = log.With().Str("module", "parser").Logger()
	parserWorkers := parser.NewWorkerPool(5, filesPipeline.Channel, resultsPipeline.Channel)
	go parserWorkers.Run()

	//_, err := parser.ParseAls(".samples/sample-001-v11-empty.als")
	////_, err := parser.ParseAls(".samples/800-ios-note-casolare.als")
	//if err != nil {
	//	panic(err)
	//}

	webservice.Serve(pusher, ":10000")
}
