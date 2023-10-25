package collector

import (
	"github.com/adrianrudnik/ablegram/pipeline"
	"github.com/adrianrudnik/ablegram/pusher"
	"github.com/rs/zerolog"
	"io/fs"
	"os"
	"path/filepath"
)

var Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

type Collection struct {
	files []string
}

func Collect(path string, filesChan chan<- *pipeline.FilesForProcessorMsg, broadcastChan chan<- interface{}) error {
	allowedExtensions := []string{".als"}

	err := findFilesByExtension(path, allowedExtensions, filesChan, broadcastChan)
	if err != nil {
		return err
	}

	return nil
}

func findFilesByExtension(root string, extensions []string, filesChan chan<- *pipeline.FilesForProcessorMsg, broadcastChan chan<- interface{}) error {
	err := filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
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
