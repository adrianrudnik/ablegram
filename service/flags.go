package main

import (
	"flag"
	"github.com/adrianrudnik/ablegram/internal/config"
)

func parseFlags(c *config.Config) {
	logLevel := flag.String("log-level", "info", "Set the log level [debug, info]")
	logToFiles := flag.Bool("enable-logs", false, "Enable debug log writing to files")
	logScannedFolders := flag.Bool("enable-scanned-log", false, "Enable scanned paths log file")

	demoMode := flag.Bool("demo-mode", false, "Enable demo mode, only readable actions will be executed")
	noBrowserFlag := flag.Bool("no-browser", false, "Skip the automatic browser opening")
	noGuiFlag := flag.Bool("no-gui", false, "Do no start the GUI.")
	noWebserviceFlag := flag.Bool("no-webservice", false, "Do no start the webservice")

	indexerWorkerDelay := flag.Int("indexer-worker-delay", 0, "Set the delay in milliseconds between indexer workers tasks")

	flag.Parse()

	c.Log.Level = *logLevel
	c.Log.EnableRuntimeLogfile = *logToFiles
	c.Log.EnableProcessedLogfile = *logScannedFolders

	c.Behaviour.DemoMode = *demoMode
	c.Behaviour.OpenBrowserOnStart = !*noBrowserFlag
	c.Behaviour.ShowServiceGui = !*noGuiFlag
	c.Behaviour.AutostartWebservice = !*noWebserviceFlag

	c.Indexer.WorkerDelayInMs = *indexerWorkerDelay
}
