package main

import (
	"flag"
	"github.com/adrianrudnik/ablegram/internal/config"
)

func parseFlags(c *config.Config) {
	logLevel := flag.String("log-level", "info", "Set the log level [debug, info]")
	logToFiles := flag.Bool("enable-logs", false, "Enable debug log writing to files")
	logScannedFolders := flag.Bool("enable-scanned-log", false, "Enable scanned paths log file")

	noBrowserFlag := flag.Bool("no-browser", false, "Skip the automatic browser opening")
	noGuiFlag := flag.Bool("no-gui", false, "Do no start the GUI.")
	noWebserviceFlag := flag.Bool("no-webservice", false, "Do no start the webservice")

	flag.Parse()

	c.Log.Level = *logLevel
	c.Log.ToFiles = *logToFiles
	c.Log.ScannedFolders = *logScannedFolders

	c.Behaviour.BrowserAutostart = !*noBrowserFlag
	c.Behaviour.ShowGui = !*noGuiFlag
	c.Behaviour.WebserviceAutostart = !*noWebserviceFlag
}
