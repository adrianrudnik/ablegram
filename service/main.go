package main

import (
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
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"image/color"
	"os"
	"time"
)

//go:generate fyne bundle -o bundled.go assets/icon.png
//go:generate fyne bundle -o bundled.go -append assets/logo.png

func main() {
	// Let's look for a configuration within one of the folders
	config.Logger = log.With().Str("module", "config").Logger()
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	appConfig := config.LoadWithDefaults("")

	parseFlags(appConfig)

	// Set up logging
	switch appConfig.Log.Level {
	case "debug":
		gin.SetMode(gin.DebugMode)
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		break
	default:
		gin.SetMode(gin.ReleaseMode)
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	if appConfig.Log.ToFiles {
		logPath := config.GetRelativeFilePath(".runtime.log")
		logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Error().Err(err).Msg("Could not open log file")
		} else {
			defer logFile.Close()
			log.Logger = zerolog.New(logFile).With().Timestamp().Logger()
		}
	}

	log.Info().Msg("App starting")

	// Create some channel based pipelines to pass around the different workloads
	pusherPipeline := pipeline.NewFrontendPush()
	filesPipeline := pipeline.NewFilesForProcessor()
	resultsPipeline := pipeline.NewDocumentsToIndex()

	// ProcessProgress is responsible in holding the current progress and
	// notifying the frontend about it
	stats.Logger = log.With().Str("module", "stats").Logger()
	progress := stats.NewProcessProgress(pusherPipeline.Chan)

	// Kick of the webservice
	go func() {
		if !appConfig.Behaviour.WebserviceAutostart {
			return
		}

		// Set the logger for the UI helper
		ui.Logger = log.With().Str("module", "ui").Logger()

		// Metrics is responsible in keeping and communicating key metrics for the frontend
		appMetrics := stats.NewMetrics(pusherPipeline.Chan)

		// Start the frontend push worker
		webservice.Logger = log.With().Str("module", "webservice").Logger()
		appPusher := webservice.NewPushChannel(pusherPipeline.Chan)
		go appPusher.Run()

		// Collector is responsible for finding files that could be parsed
		collector.Logger = log.With().Str("module", "collector").Logger()
		collectorWorkers := collector.NewWorkerPool(10, filesPipeline.Chan, pusherPipeline.Chan)
		go collectorWorkers.Run(appConfig, progress)

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
			if !appConfig.Behaviour.BrowserAutostart {
				return
			}

			time.Sleep(50 * time.Millisecond)
			ui.OpenFrontend()
		}()

		webservice.Serve(search, appPusher, ":10000")
	}()

	if !appConfig.Behaviour.ShowGui {
		select {}
	} else {
		// Define a clean theme

		a := app.New()
		a.Settings().SetTheme(&ui.AblegramTheme{})
		a.SetIcon(resourceIconPng)
		w := a.NewWindow("Ablegram")
		w.CenterOnScreen()

		logo := canvas.NewImageFromResource(resourceLogoPng)
		logo.FillMode = canvas.ImageFillOriginal

		statusTxt := canvas.NewText("The service is processing files...", color.White)
		quitBtn := widget.NewButton("Shut down service", func() { a.Quit() })
		startBtn := widget.NewButton("Open results in browser", func() { ui.OpenFrontend() })
		progressBar := widget.NewProgressBarInfinite()

		uiUpdater := ui.NewUiUpdater(statusTxt, progressBar)
		go uiUpdater.Run(progress)

		content := container.New(layout.NewPaddedLayout(), container.New(layout.NewVBoxLayout(), logo, statusTxt, progressBar, startBtn, quitBtn))

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
