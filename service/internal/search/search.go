package search

import (
	"github.com/blevesearch/bleve/analysis/lang/en"
	"github.com/blevesearch/bleve/v2"
	"github.com/rs/zerolog"
	"os"
)

var Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

type Search struct {
	options *SearchOptions

	Index bleve.Index
}

func NewSearch(options *SearchOptions) *Search {
	indexMapping := bleve.NewIndexMapping()
	index, err := bleve.NewMemOnly(indexMapping)
	if err != nil {
		Logger.Panic().Err(err).Msg("Failed to create memory index")
		panic(err)
	}

	indexMapping.AddDocumentMapping("als_file", buildAlsFileMapping(options))
	indexMapping.AddDocumentMapping("audio_track", buildAudioTrackMapping(options))
	indexMapping.AddDocumentMapping("midi_track", buildMidiTrackMapping(options))

	indexMapping.DefaultAnalyzer = en.AnalyzerName

	Logger.Info().Msg("Indexes created")

	return &Search{
		options: options,
		Index:   index,
	}
}
