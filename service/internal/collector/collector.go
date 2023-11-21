package collector

import (
	"github.com/adrianrudnik/ablegram/internal/config"
	"github.com/adrianrudnik/ablegram/internal/pusher"
	"github.com/adrianrudnik/ablegram/internal/workload"
	"github.com/rs/zerolog"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

var Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

var excludePaths = []string{
	os.TempDir(),
}

var excludeFolders = []string{
	"node_modules",
}

var allowedExtensions = []string{
	".als",
}

type Collection struct {
	files []string
}

func Collect(
	conf *config.Config,
	target *config.CollectorTarget,

	filesChan chan<- *workload.FilePayload,
	broadcastChan chan<- interface{},
) error {
	err := findFilesByExtension(conf, target, allowedExtensions, filesChan, broadcastChan)
	if err != nil {
		return err
	}

	return nil
}

func findFilesByExtension(
	conf *config.Config,
	target *config.CollectorTarget,
	extensions []string,
	fileChan chan<- *workload.FilePayload,
	pushChan chan<- interface{},
) error {

	folders := make([]string, 0, 1000000)

	// Exclude OS specific folders
	enrichPlatformExcludePaths()

	err := filepath.WalkDir(target.Uri, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			Logger.Warn().Err(e).Str("path", s).Msg("Skipped folder due to error")
			return nil
		}

		// Exclude folders beginning with a dot
		if target.ExcludeDotFolders && d.IsDir() && strings.HasPrefix(d.Name(), ".") {
			Logger.Debug().Str("path", s).Msg("Skipping dot folder")
			return filepath.SkipDir
		}

		// Exclude paths by prefix
		if target.ExcludeSystemFolders && d.IsDir() && slices.IndexFunc(excludePaths, func(f string) bool {
			return strings.HasPrefix(s, f)
		}) != -1 {
			Logger.Debug().Str("path", s).Msg("Skipping excluded system path")
			return filepath.SkipDir
		}

		// Exclude folders by name
		if target.ExcludeSystemFolders && d.IsDir() && slices.IndexFunc(excludeFolders, func(s string) bool {
			return s == filepath.Base(d.Name())
		}) != -1 {
			Logger.Debug().Str("path", s).Msg("Skipping excluded system folder")
			return filepath.SkipDir
		}

		// Collect the folders we visited for the processed log file
		if conf.Log.EnableProcessedLogfile && d.IsDir() {
			folders = append(folders, s)
		}

		for _, ext := range extensions {
			if !d.IsDir() && filepath.Ext(d.Name()) == ext {
				Logger.Debug().Str("file", s).Msg("Found file for processing")

				// Notify the UI about this file
				pushChan <- pusher.NewFileStatusPush(s, "pending", "")

				// Move it over to the processing pipeline
				fileChan <- workload.NewFilePayload(s)
			}
		}

		return nil
	})

	if conf.Log.EnableProcessedLogfile && len(folders) > 0 {
		scanLogPath := config.GetRelativeFilePath(".scanned-folders.log")
		lines := strings.Join(folders, "\n")
		err := os.WriteFile(scanLogPath, []byte(lines), 0666)
		if err != nil {
			Logger.Warn().Err(err).Msg("Failed to write scanned folders to file")
		}
	}

	if err != nil {
		return err
	}

	return nil
}
