package collector

import (
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
	".git",
	".idea",
}

type Collection struct {
	files []string
}

func Collect(path string, filesChan chan<- *pipeline.FilesForProcessorMsg, broadcastChan chan<- interface{}) error {
	wg.Add(1)
	defer wg.Done()

	broadcastChan <- pusher.NewProcessingStatusPush(true)

	allowedExtensions := []string{".als"}

	err := findFilesByExtension(path, allowedExtensions, filesChan, broadcastChan)
	if err != nil {
		return err
	}

	go func() {
		wg.Wait()
		broadcastChan <- pusher.NewProcessingStatusPush(false)
	}()

	return nil
}

func findFilesByExtension(root string, extensions []string, filesChan chan<- *pipeline.FilesForProcessorMsg, broadcastChan chan<- interface{}) error {
	err := filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		// Exclude paths by prefix
		if d.IsDir() && slices.IndexFunc(excludePaths, func(s string) bool {
			return strings.HasPrefix(d.Name(), s)
		}) != -1 {
			Logger.Info().Str("path", s).Msg("Skipping excluded path")
			return filepath.SkipDir
		}

		// Exclude folders by name
		if d.IsDir() && slices.IndexFunc(excludeFolders, func(s string) bool {
			return s == filepath.Base(d.Name())
		}) != -1 {
			Logger.Info().Str("path", s).Msg("Skipping excluded folder")
			return filepath.SkipDir
		}

		if e != nil {
			Logger.Warn().Err(e).Str("path", s).Msg("Skipped folder due to error")
			return nil
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

	if err != nil {
		return err
	}

	return nil
}
