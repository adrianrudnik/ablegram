package main

import (
	"flag"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/adrianrudnik/ablegram/internal/collector"
	"github.com/adrianrudnik/ablegram/internal/config"
	"github.com/adrianrudnik/ablegram/internal/indexer"
	"github.com/adrianrudnik/ablegram/internal/parser"
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/pusher"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/ui"
	"github.com/adrianrudnik/ablegram/internal/webservice"
	"github.com/icza/gox/osx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"image/color"
	"os"
	"time"
)

//go:generate fyne bundle -o bundled.go assets/icon.png
//go:generate fyne bundle -o bundled.go -append assets/logo.png

func main() {
	// Parse flags
	noBrowserFlag := flag.Bool("no-browser", false, "Skip the automatic browser opening")
	noGuiFlag := flag.Bool("no-gui", false, "Do no start the GUI")
	noWebserviceFlag := flag.Bool("no-webservice", false, "Do no start the webservice")

	flag.Parse()

	log.Info().
		Bool("no-browser", !*noBrowserFlag).
		Bool("no-gui", !*noGuiFlag).
		Bool("no-webservice", !*noWebserviceFlag).
		Msg("Parsed executable flags")

	// Set up logging
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("App starting")

	// Create some channel based pipelines to pass around the different workloads
	pusherPipeline := pipeline.NewFrontendPush()
	filesPipeline := pipeline.NewFilesForProcessor()
	resultsPipeline := pipeline.NewDocumentsToIndex()

	// ProcessProgress is responsible in holding the current progress and
	// notifying the frontend about it
	progress := stats.NewProcessProgress(pusherPipeline.Chan)

	// Kick of the webservice
	go func() {
		if *noWebserviceFlag {
			return
		}

		// Let's look for a configuration within one of the folders
		config.Logger = log.With().Str("module", "config").Logger()
		appConfig := config.LoadWithDefaults("")

		// Metrics is responsible in keeping and communicating key metrics for the frontend
		appMetrics := stats.NewMetrics(pusherPipeline.Chan)

		// Start the frontend push worker
		webservice.Logger = log.With().Str("module", "webservice").Logger()
		appPusher := webservice.NewPushChannel(pusherPipeline.Chan)
		go appPusher.Run()

		// Collector is responsible for finding files that could be parsed
		collector.Logger = log.With().Str("module", "collector").Logger()
		collectorWorkers := collector.NewWorkerPool(10, filesPipeline.Chan, pusherPipeline.Chan)
		go collectorWorkers.Run(progress, appConfig.SearchablePaths)

		// Parser is responsible for parsing the files into results for the indexerWorker
		parser.Logger = log.With().Str("module", "parser").Logger()
		parserWorkers := parser.NewWorkerPool(10, filesPipeline.Chan, resultsPipeline.Chan, pusherPipeline.Chan)
		go parserWorkers.Run(progress, appMetrics)

		// Create the indexerWorker
		indexer.Logger = log.With().Str("module", "indexer").Logger()
		search := indexer.NewSearch()
		indexerWorker := indexer.NewWorker(search, resultsPipeline.Chan, pusherPipeline.Chan)
		go indexerWorker.Run(progress, appMetrics)

		// Try to open the default browser on the given OS
		go func() {
			if *noBrowserFlag {
				return
			}

			time.Sleep(50 * time.Millisecond)
			openBrowser()
		}()

		webservice.Serve(search, appPusher, ":10000")
	}()

	if *noGuiFlag {
		select {}
	} else {
		// Define a clean theme

		a := app.New()
		a.Settings().SetTheme(&ablegramTheme{})
		a.SetIcon(resourceIconPng)
		w := a.NewWindow("Ablegram")
		w.CenterOnScreen()

		logo := canvas.NewImageFromResource(resourceLogoPng)
		logo.FillMode = canvas.ImageFillOriginal

		statusTxt := canvas.NewText("Starting up...", color.White)
		quitBtn := widget.NewButton("Shut down service", func() { a.Quit() })
		startBtn := widget.NewButton("Open results in browser", func() { openBrowser() })

		uiUpdater := ui.NewUiUpdater(statusTxt)
		go uiUpdater.Run(progress)

		content := container.New(layout.NewPaddedLayout(), container.New(layout.NewVBoxLayout(), logo, statusTxt, startBtn, quitBtn))

		w.SetContent(content)
		w.SetFixedSize(true)

		// Hotfix for https://github.com/fyne-io/fyne/issues/4350
		w.Resize(fyne.NewSize(1, 1))

		w.ShowAndRun()

		// Say goodbye in the browser window, if available.
		pusherPipeline.Chan <- pusher.NewNavigatePush("goodbye")
		time.Sleep(100 * time.Millisecond)
	}
}

func openBrowser() {
	err := osx.OpenDefault("http://localhost:10000")
	if err != nil {
		log.Warn().Err(err).Msg("Could not open default browser")
	}
}
