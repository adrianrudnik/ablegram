package collector

import (
	"github.com/adrianrudnik/ablegram/internal/config"
	"github.com/adrianrudnik/ablegram/internal/pipeline"
	"github.com/adrianrudnik/ablegram/internal/pusher"
	"github.com/rs/zerolog"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
)

var Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

var wg sync.WaitGroup

var excludePaths = []string{
	os.TempDir(),
}

var excludeFolders = []string{
	"node_modules",
}

type Collection struct {
	files []string
}

func Collect(c *config.Config, path string, filesChan chan<- *pipeline.FilesForProcessorMsg, broadcastChan chan<- interface{}) error {
	allowedExtensions := []string{".als"}

	err := findFilesByExtension(c, path, allowedExtensions, filesChan, broadcastChan)
	if err != nil {
		return err
	}

	go func() {
		wg.Wait()
	}()

	return nil
}

func findFilesByExtension(c *config.Config, root string, extensions []string, filesChan chan<- *pipeline.FilesForProcessorMsg, broadcastChan chan<- interface{}) error {

	folders := make([]string, 0, 1000000)

	err := filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			Logger.Warn().Err(e).Str("path", s).Msg("Skipped folder due to error")
			return nil
		}

		// Exclude folders beginning with a dot
		if c.Collector.ExcludeSystemFolders && d.IsDir() && strings.HasPrefix(d.Name(), ".") {
			Logger.Debug().Str("path", s).Msg("Skipping dot folder")
			return filepath.SkipDir
		}

		// Exclude paths by prefix
		if c.Collector.ExcludeSystemFolders && d.IsDir() && slices.IndexFunc(excludePaths, func(s string) bool {
			return strings.HasPrefix(d.Name(), s)
		}) != -1 {
			Logger.Debug().Str("path", s).Msg("Skipping excluded path")
			return filepath.SkipDir
		}

		// Exclude folders by name
		if c.Collector.ExcludeSystemFolders && d.IsDir() && slices.IndexFunc(excludeFolders, func(s string) bool {
			return s == filepath.Base(d.Name())
		}) != -1 {
			Logger.Debug().Str("path", s).Msg("Skipping excluded folder")
			return filepath.SkipDir
		}

		// Log at least the folders we are visiting
		if d.IsDir() {
			folders = append(folders, s)
		}

		for _, ext := range extensions {
			if !d.IsDir() && filepath.Ext(d.Name()) == ext {
				Logger.Debug().Str("file", s).Msg("Found possible file")

				// Notify the UI about this file
				broadcastChan <- pusher.NewFileStatusPush(s, "pending", "")

				// Move it over to the processing pipeline
				filesChan <- &pipeline.FilesForProcessorMsg{AbsPath: s}
			}
		}

		return nil
	})

	if c.Log.ScannedFolders {
		scanLogPath := config.GetRelativeFilePath(".scanned-folders.log")
		lines := strings.Join([]string(folders), "\n")
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
