package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/adrianrudnik/ablegram/internal/access"
	"github.com/adrianrudnik/ablegram/internal/collector"
	"github.com/adrianrudnik/ablegram/internal/config"
	"github.com/adrianrudnik/ablegram/internal/indexer"
	"github.com/adrianrudnik/ablegram/internal/parser"
	"github.com/adrianrudnik/ablegram/internal/pushermsg"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"github.com/adrianrudnik/ablegram/internal/suggest"
	"github.com/adrianrudnik/ablegram/internal/tagger"
	"github.com/adrianrudnik/ablegram/internal/ui"
	"github.com/adrianrudnik/ablegram/internal/webservice"
	"github.com/adrianrudnik/ablegram/internal/workload"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
	"time"
)

// Generate the bundled.go file containing all static assets used by the GUI
//go:generate fyne bundle -o bundled.go assets/icon.png
//go:generate fyne bundle -o bundled.go -append assets/logo-wide-light.static.png
//go:generate fyne bundle -o bundled.go -append assets/logo-wide-dark.static.png

var AppVersion = "dev"
var BuildCommit = "unknown"
var BuildDate = "unknown"

func main() {
	// Let's look for a configuration within one of the folders
	config.Logger = log.With().Str("module", "config").Logger()
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	// Set up the app configuration with defaults for now
	appConfig := config.LoadWithDefaults("")
	appConfig.About.Version = AppVersion
	appConfig.About.Commit = BuildCommit
	appConfig.About.Date = BuildDate

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

	if appConfig.Log.EnableRuntimeLogfile {
		logPath := config.GetRelativeFilePath(".runtime.log")
		logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Error().Err(err).Msg("Could not open log file")
		} else {
			defer logFile.Close()
			log.Logger = zerolog.New(logFile).With().Timestamp().Logger()
		}
	}

	log.Info().
		Str("build-version", AppVersion).
		Str("build-commit", BuildCommit).
		Str("build-date", BuildDate).
		Msg("App starting")

	// Channel to push messages to the frontend
	pushChan := make(chan workload.PushMessage, 5000)

	// Channel to inject parsed results to be indexed by the index workers
	// We go for 20k buffer, as we will have at least 10 documents per parsed file
	// but the indexer can only batch single threaded.
	resultsToIndexChan := make(chan *workload.DocumentPayload, 50000)

	// ProcessProgress is responsible in holding the current appProgress and
	// notifying the frontend about it
	stats.Logger = log.With().Str("module", "stats").Logger()
	appProgress := stats.NewProcessProgress(pushChan)

	// TagCollector is responsible for collecting all tags and pushing them to the frontend
	// if the collector is wired to a push channel
	tagger.Logger = log.With().Str("module", "tagger").Logger()
	appTags := tagger.NewTagCollector()
	appTags.WirePusher(pushChan)

	// Start up the auth and otp services
	appOtp := access.NewOtp()
	appAuth := access.NewAuth(appOtp)
	appUsers := access.NewUserList()

	// Start the suggestion service that allows guests to suggest stuff to admins
	appSuggest := suggest.NewList()

	// Kick of the webservice
	go func() {
		if !appConfig.Behaviour.AutostartWebservice {
			return
		}

		// Set the logger for the UI helper
		ui.Logger = log.With().Str("module", "ui").Logger()

		// Statistics is responsible in keeping and communicating key metrics for the frontend
		appStats := stats.NewStatistics(appConfig, pushChan)

		// Start the frontend push worker
		webservice.Logger = log.With().Str("module", "webservice").Logger()

		parser.Logger = log.With().Str("module", "parser").Logger()

		// Collector is responsible for finding files that could be parsed
		collector.Logger = log.With().Str("module", "collector").Logger()
		collectorWorkers := collector.NewWorkerPool(appConfig, appStats, appProgress, appTags, resultsToIndexChan, pushChan)
		go collectorWorkers.Run()

		// Create the indexerWorker
		indexer.Logger = log.With().Str("module", "indexer").Logger()
		appIndexer := indexer.NewSearch()
		indexerWorker := indexer.NewWorker(appConfig, appIndexer, resultsToIndexChan, pushChan)
		go indexerWorker.Run(appProgress, appStats)

		// Try to open the default browser on the given OS
		go func() {
			if !appConfig.Behaviour.OpenBrowserOnStart {
				return
			}

			time.Sleep(50 * time.Millisecond)
			ui.OpenFrontendAsAdmin(appConfig, appOtp)
		}()

		for _, port := range appConfig.Webservice.TryPorts {
			appConfig.Webservice.ChosenPort = port
			err := webservice.Serve(
				appConfig,
				appAuth,
				appOtp,
				appUsers,
				appIndexer,
				appTags,
				appSuggest,
				pushChan,
				fmt.Sprintf(":%d", port),
			)
			if err != nil && strings.Contains(err.Error(), "bind: permission denied") {
				log.Warn().Err(err).Int("port", port).Msg("Could not start webservice, trying other port")
				continue
			}
			break
		}
	}()

	if !appConfig.Behaviour.ShowServiceGui {
		select {}
	} else {
		// Define a clean theme

		a := app.New()
		a.SetIcon(resourceIconPng)
		a.Settings().SetTheme(&ui.AblegramTheme{})

		w := a.NewWindow("Ablegram")
		w.CenterOnScreen()

		var logo *canvas.Image
		if a.Settings().ThemeVariant() == theme.VariantLight {
			log.Debug().Msg("UI is using light theme")
			logo = canvas.NewImageFromResource(resourceLogoWideLightStaticPng)
		} else {
			log.Debug().Msg("UI is using dark theme")
			logo = canvas.NewImageFromResource(resourceLogoWideDarkStaticPng)
		}
		logo.FillMode = canvas.ImageFillOriginal

		statusTxt := canvas.NewText("The service is processing files...", theme.ForegroundColor())
		quitBtn := widget.NewButton("Shutdown", func() { a.Quit() })
		startBtn := widget.NewButton("Open search", func() { ui.OpenFrontendAsAdmin(appConfig, appOtp) })

		progressBar := widget.NewProgressBarInfinite()

		// @see https://github.com/fyne-io/fyne/issues/2469#issuecomment-1789642706
		quitBtn.Importance = widget.HighImportance
		startBtn.Importance = widget.HighImportance

		uiUpdater := ui.NewUiUpdater(statusTxt, progressBar)
		go uiUpdater.Run(appProgress)

		content := container.New(layout.NewPaddedLayout(), container.New(layout.NewVBoxLayout(), logo, statusTxt, progressBar, startBtn, quitBtn))

		w.SetContent(content)
		w.SetFixedSize(true)

		// Hotfix for https://github.com/fyne-io/fyne/issues/4350
		w.Resize(fyne.NewSize(1, 1))

		w.ShowAndRun()

		// Say goodbye in the browser window, if available.
		pushChan <- pushermsg.NewForceNavigatePush("goodbye")
		time.Sleep(100 * time.Millisecond)
	}
}
